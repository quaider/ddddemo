package cargo

import (
	"errors"
	"go-ddd/domain/location"
	"time"
)

// RouteSpecification 表示 Cargo 的起始地和目的地、截止到达时间
type RouteSpecification struct {
	origin          *location.Location
	destination     *location.Location
	arrivalDeadline time.Time
}

func NewRouteSpecification(origin *location.Location, destination *location.Location, arrivalDeadline time.Time) (*RouteSpecification, error) {
	if origin == nil {
		return nil, errors.New("origin is required")
	}

	if destination == nil {
		return nil, errors.New("destination is required")
	}

	if origin == destination {
		return nil, errors.New("origin and destination can't be same")
	}

	return &RouteSpecification{
		origin:          origin,
		destination:     destination,
		arrivalDeadline: arrivalDeadline,
	}, nil
}

func (r *RouteSpecification) Origin() *location.Location {
	return r.origin
}

func (r *RouteSpecification) Destination() *location.Location {
	return r.destination
}

func (r *RouteSpecification) ArrivalDeadline() time.Time {
	return r.arrivalDeadline
}

func (r *RouteSpecification) isSatisfiedBy(itinerary *Itinerary) bool {
	return itinerary != nil &&
		r.origin.SameIdentityAs(itinerary.InitialDepartureLocation()) &&
		r.destination.SameIdentityAs(itinerary.FinalArrivalLocation()) &&
		r.arrivalDeadline.After(itinerary.FinalArrivalDate())

}
