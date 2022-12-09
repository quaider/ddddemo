package handling

import (
	"errors"
	"go-ddd/domain/location"
	"go-ddd/domain/voyage"
	"time"
)

func (e EventType) RequiresVoyage() bool {
	return bool(e)
}

func (e EventType) ProhibitsVoyage() bool {
	return !e.RequiresVoyage()
}

// Event 用于注册事件，例如，在给定时间某个位置从承运人卸载货物 Cargo 时。
// 此类是HandlingEvent聚合的唯一成员，因此也是根
type Event struct {
	id               int // 自动生成，暂时忽略
	eventType        EventType
	voyage           *voyage.Voyage
	location         *location.Location
	completionTime   time.Time // 事件完成事件
	registrationTime time.Time // 事件注册事件
	cargoTrackingId  string    // 关联的货物id (直接引用cargo会导致 import cycle)
}

func NewEventWithoutVoyage(eventType EventType, location *location.Location, completionTime time.Time, registrationTime time.Time, cargoTrackingId string) (*Event, error) {
	if cargoTrackingId == "" {
		return nil, errors.New("cargoTrackingId is required")
	}

	if location == nil {
		return nil, errors.New("location is required")
	}

	if eventType.RequiresVoyage() {
		return nil, errors.New("voyage is required for event type")
	}

	return &Event{
		eventType:        eventType,
		location:         location,
		completionTime:   completionTime,
		registrationTime: registrationTime,
		cargoTrackingId:  cargoTrackingId,
	}, nil
}

func NewEvent(eventType EventType, voyage *voyage.Voyage, location *location.Location, completionTime, registrationTime time.Time, cargoTrackingId string) (*Event, error) {
	if cargoTrackingId == "" {
		return nil, errors.New("cargoTrackingId is required")
	}

	if location == nil {
		return nil, errors.New("location is required")
	}

	if voyage == nil {
		return nil, errors.New("voyage is required")
	}

	if eventType.ProhibitsVoyage() {
		return nil, errors.New("voyage is not allowed with event type")
	}

	return &Event{
		eventType:        eventType,
		voyage:           voyage,
		location:         location,
		completionTime:   completionTime,
		registrationTime: registrationTime,
		cargoTrackingId:  cargoTrackingId,
	}, nil
}

func (e *Event) Id() int {
	return e.id
}

func (e *Event) EventType() EventType {
	return e.eventType
}

func (e *Event) Voyage() *voyage.Voyage {
	return e.voyage
}

func (e *Event) Location() *location.Location {
	return e.location
}

func (e *Event) CompletionTime() time.Time {
	return e.completionTime
}

func (e *Event) RegistrationTime() time.Time {
	return e.registrationTime
}

func (e *Event) CargoTrackingId() string {
	return e.cargoTrackingId
}
