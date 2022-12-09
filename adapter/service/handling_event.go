package service

import (
	"go-ddd/domain"
	"go-ddd/domain/handling"
	"go-ddd/domain/location"
	"go-ddd/domain/voyage"
	"log"
	"time"
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
	locations location.Repository
	voyages   voyage.Repository
	events    handling.Repository
}

func NewHandlingEventServiceImpl(locations location.Repository, voyages voyage.Repository, events handling.Repository) HandlingEventService {
	return &handlingEventServiceImpl{locations: locations, voyages: voyages, events: events}
}

func (h *handlingEventServiceImpl) RegisterHandlingEvent(
	completionTime time.Time,
	trackingId domain.TrackingId,
	voyageNumber string,
	unlocode string,
	eventType handling.EventType) error {

	registrationTime := time.Now()
	event, err := handling.NewEventFactory(h.voyages, h.locations).
		CreateHandlingEvent(registrationTime, completionTime, trackingId, voyageNumber, unlocode, eventType)
	if err != nil {
		return err
	}

	// save event
	err = h.events.Save(event)
	if err != nil {
		return err
	}

	// publish an event stating that a cargo has been handled

	log.Println("registered handling event")

	return nil
}
