package cargo

import (
	"errors"
	"go-ddd/domain/handling"
	"go-ddd/domain/location"
	"go-ddd/domain/voyage"
	"time"
)

type Delivery struct {
	transportStatus         TransportStatus
	lastLocation            *location.Location
	currentVoyage           *voyage.Voyage
	misdirected             bool
	eta                     time.Time //  Estimated time of arrival 预估到达时间
	nextExpectedActivity    *HandlingActivity
	isUnloadedAtDestination bool
	routingStatus           RoutingStatus
	calculatedAt            time.Time
	lastEvent               *handling.Event
}

func newDelivery(lastEvent *handling.Event, itinerary *Itinerary, route *RouteSpecification) *Delivery {
	d := &Delivery{
		calculatedAt: time.Now(),
		lastEvent:    lastEvent,
	}

	d.misdirected = d.calculateMisdirectionStatus(itinerary)
	d.routingStatus = d.calculateRoutingStatus(itinerary, route)

	return d
}

func (d *Delivery) calculateMisdirectionStatus(itinerary *Itinerary) bool {
	if d.lastEvent == nil {
		return false
	}

	return itinerary.IsExpected(d.lastEvent)
}

// UpdateOnRouting 创建 Delivery 快照 以应对路径或航线的变更
func (d *Delivery) UpdateOnRouting(route *RouteSpecification, itinerary *Itinerary) (*Delivery, error) {
	if route == nil {
		return nil, errors.New("route specification is required")
	}

	return newDelivery(d.lastEvent, itinerary, route), nil
}

func (d *Delivery) calculateRoutingStatus(itinerary *Itinerary, route *RouteSpecification) RoutingStatus {
	if itinerary == nil {
		return NotRouted
	}

	if route.isSatisfiedBy(itinerary) {
		return Routed
	}

	return MisRouted
}

func DerivedFrom(route *RouteSpecification, it *Itinerary, handlingHistory *handling.History) (*Delivery, error) {
	if route == nil {
		return nil, errors.New("route specification is required")
	}

	// 从历史中获取 最后 一个 handling event
	lastEvent := handlingHistory.MostRecentlyCompletedEvent()

	return newDelivery(lastEvent, it, route), nil
}
