package service

import (
	"go-ddd/domain"
	"go-ddd/domain/handling"

	"log"
	"time"

	"github.com/asaskevich/EventBus"
)

var _ HandlingEventService = (*handlingEventServiceImpl)(nil)

type HandlingEventService interface {
	RegisterHandlingEvent(
		completionTime time.Time,
		trackingId domain.TrackingId,
		voyageNumber string,
		unlocode string,
		eventType handling.EventType) error
}

type handlingEventServiceImpl struct {
	events       handling.Repository
	eventFactory *handling.EventFactory
	bus          EventBus.Bus // 应该自己封装一套接口，与EventBus的实现隔离
}

func NewHandlingEventServiceImpl(events handling.Repository, eventFactory *handling.EventFactory, bus EventBus.Bus) HandlingEventService {
	return &handlingEventServiceImpl{events: events, eventFactory: eventFactory, bus: bus}
}

func (h *handlingEventServiceImpl) RegisterHandlingEvent(
	completionTime time.Time,
	trackingId domain.TrackingId,
	voyageNumber string,
	unlocode string,
	eventType handling.EventType) error {

	registrationTime := time.Now()
	event, err := h.eventFactory.CreateHandlingEvent(registrationTime, completionTime, trackingId, voyageNumber, unlocode, eventType)
	if err != nil {
		return err
	}

	// save event
	err = h.events.Save(event)
	if err != nil {
		return err
	}

	// publish an event stating that a cargo has been handled
	h.bus.Publish("cargo:cargoHandled", event)

	log.Println("registered handling event")

	return nil
}
