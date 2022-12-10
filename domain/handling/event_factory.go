package handling

import (
	"errors"
	"go-ddd/domain"
	"go-ddd/domain/location"
	"go-ddd/domain/voyage"
	"time"
)

type EventFactory struct {
	voyageRepository   voyage.Repository   `container:"type"`
	locationRepository location.Repository `container:"type"`
}

func init() {

}

func NewEventFactory(voyageRepository voyage.Repository, locationRepository location.Repository) *EventFactory {
	return &EventFactory{voyageRepository: voyageRepository, locationRepository: locationRepository}
}

func (factory *EventFactory) CreateHandlingEvent(
	registrationTime, completionTime time.Time,
	trackingId domain.TrackingId, voyageNumber, unlocode string, eventType EventType) (*Event, error) {
	l := factory.locationRepository.FindLocation(unlocode)
	if l == nil {
		return nil, errors.New("unknow location with " + unlocode)
	}

	if voyageNumber != "" {
		return NewEventWithoutVoyage(eventType, l, completionTime, registrationTime, trackingId)
	}

	v := factory.voyageRepository.FindVoyage(voyageNumber)
	if v == nil {
		return NewEventWithoutVoyage(eventType, l, completionTime, registrationTime, trackingId)
	}

	return NewEvent(eventType, v, l, completionTime, registrationTime, trackingId)
}
