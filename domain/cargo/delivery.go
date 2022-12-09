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

func (d *Delivery) TransportStatus() TransportStatus {
	return d.transportStatus
}

func (d *Delivery) LastLocation() *location.Location {
	return d.lastLocation
}

func (d *Delivery) CurrentVoyage() *voyage.Voyage {
	return d.currentVoyage
}

func (d *Delivery) Misdirected() bool {
	return d.misdirected
}

func (d *Delivery) Eta() time.Time {
	return d.eta
}

func (d *Delivery) NextExpectedActivity() *HandlingActivity {
	return d.nextExpectedActivity
}

func (d *Delivery) IsUnloadedAtDestination() bool {
	return d.isUnloadedAtDestination
}

func (d *Delivery) RoutingStatus() RoutingStatus {
	return d.routingStatus
}

func (d *Delivery) CalculatedAt() time.Time {
	return d.calculatedAt
}

func (d *Delivery) LastEvent() *handling.Event {
	return d.lastEvent
}

func newDelivery(lastEvent *handling.Event, itinerary *Itinerary, route *RouteSpecification) *Delivery {
	d := &Delivery{
		calculatedAt: time.Now(),
		lastEvent:    lastEvent,
	}

	d.misdirected = d.calculateMisdirectionStatus(itinerary)
	d.routingStatus = d.calculateRoutingStatus(itinerary, route)
	d.transportStatus = calculateTransportStatus(d)
	d.lastLocation = calculateLastKnownLocation(d)
	d.currentVoyage = calculateCurrentVoyage(d)
	d.eta = d.calculateEta(itinerary)
	activity, err := d.calculateNextExpectedActivity(route, itinerary)
	if err == nil {
		d.nextExpectedActivity = activity
	}

	d.isUnloadedAtDestination = d.calculateUnloadedAtDestination(route)

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

func calculateTransportStatus(delivery *Delivery) TransportStatus {
	if delivery.lastEvent == nil {
		return NOT_RECEIVED
	}

	switch delivery.lastEvent.EventType() {
	case handling.LOAD:
		return ONBOARD_CARRIER
	case handling.UNLOAD:
	case handling.RECEIVE:
	case handling.CUSTOMS:
		return IN_PORT
	case handling.CLAIM:
		return CLAIMED
	}

	return UNKNOWN
}

func calculateLastKnownLocation(delivery *Delivery) *location.Location {
	if delivery.lastEvent != nil {
		return delivery.lastEvent.Location()
	}

	return nil
}

func (d *Delivery) onTrack() bool {
	return d.routingStatus == Routed && !d.misdirected
}

func calculateCurrentVoyage(delivery *Delivery) *voyage.Voyage {
	if delivery.transportStatus == ONBOARD_CARRIER && delivery.lastEvent != nil {
		return delivery.lastEvent.Voyage()
	}

	return nil
}

func (d *Delivery) calculateEta(itinerary *Itinerary) time.Time {
	if d.onTrack() {
		return itinerary.FinalArrivalDate()
	}

	return time.Time{}
}

func (d *Delivery) calculateNextExpectedActivity(rs *RouteSpecification, itinerary *Itinerary) (*HandlingActivity, error) {
	if !d.onTrack() {
		return nil, nil
	}

	if d.lastEvent == nil {
		return NewHandlingActivity(handling.RECEIVE, rs.origin, nil)
	}

	switch d.lastEvent.EventType() {
	case handling.LOAD:
		for _, l := range itinerary.legs {
			if l.loadLocation.SameIdentityAs(d.lastEvent.Location()) {
				return NewHandlingActivity(handling.UNLOAD, l.unloadLocation, l.voyage)
			}
		}

		return nil, nil
	case handling.UNLOAD:
		length := len(itinerary.legs)
		for i := 0; i < length; i++ {
			if itinerary.legs[i].unloadLocation.SameIdentityAs(d.lastEvent.Location()) {
				if i < length-1 {
					return NewHandlingActivity(handling.LOAD, itinerary.legs[i+1].loadLocation, itinerary.legs[i+1].voyage)
				}

				return NewHandlingActivity(handling.CLAIM, itinerary.legs[i+1].unloadLocation, nil)
			}
		}
	case handling.RECEIVE:
		first := itinerary.legs[0]
		return NewHandlingActivity(handling.LOAD, first.loadLocation, first.voyage)
	}

	return nil, nil
}

func (d *Delivery) calculateUnloadedAtDestination(rs *RouteSpecification) bool {
	return d.lastEvent != nil &&
		handling.UNLOAD == d.lastEvent.EventType() &&
		rs.destination.SameIdentityAs(d.lastEvent.Location())
}

func DerivedFrom(route *RouteSpecification, it *Itinerary, handlingHistory *handling.History) (*Delivery, error) {
	if route == nil {
		return nil, errors.New("route specification is required")
	}

	// 从历史中获取 最后 一个 handling event
	lastEvent := handlingHistory.MostRecentlyCompletedEvent()

	return newDelivery(lastEvent, it, route), nil
}
