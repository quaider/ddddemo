package cargo

import (
	"errors"
	"go-ddd/domain/handling"
	"go-ddd/domain/location"
	"time"
)

type Itinerary struct {
	legs []*Leg
}

func NewItineraryEmpty() *Itinerary {
	return &Itinerary{
		legs: make([]*Leg, 0),
	}
}

func NewItinerary(legs []*Leg) (*Itinerary, error) {
	if legs == nil || len(legs) == 0 {
		return nil, errors.New("leg is empty")
	}

	return &Itinerary{legs: legs}, nil
}

func (i *Itinerary) lastLeg() *Leg {
	if len(i.legs) == 0 {
		return nil
	}

	return i.legs[len(i.legs)-1]
}

func (i *Itinerary) Legs() []*Leg {
	return i.legs
}

func (i *Itinerary) InitialDepartureLocation() *location.Location {
	if len(i.legs) == 0 {
		return nil
	}

	return i.legs[0].loadLocation
}

func (i *Itinerary) FinalArrivalLocation() *location.Location {
	if len(i.legs) == 0 {
		return nil
	}

	return i.lastLeg().unloadLocation
}

func (i *Itinerary) FinalArrivalDate() time.Time {
	leg := i.lastLeg()
	if leg == nil {
		return time.Now().Add(1e6 * time.Hour)
	}

	return i.lastLeg().unloadTime
}

// IsExpected Test if the given handling event is expected when executing this itinerary.
func (i *Itinerary) IsExpected(event *handling.Event) bool {
	if len(i.legs) == 0 {
		return true
	}

	if event.EventType() == handling.RECEIVE {
		return i.legs[0].LoadLocation().SameIdentityAs(event.Location())
	}

	if event.EventType() == handling.LOAD {
		for _, l := range i.legs {
			if l.LoadLocation().SameIdentityAs(event.Location()) && l.Voyage().SameIdentityAs(event.Voyage()) {
				return true
			}
		}

		return false
	}

	if event.EventType() == handling.UNLOAD {
		for _, l := range i.legs {
			if l.UnloadLocation() == event.Location() && l.Voyage() == event.Voyage() {
				return true
			}
		}

		return false
	}

	if event.EventType() == handling.CLAIM {
		leg := i.lastLeg()
		return leg.unloadLocation == event.Location()
	}

	return false
}
