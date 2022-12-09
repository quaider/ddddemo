package cargo

import (
	"errors"
	"go-ddd/domain/handling"
	"go-ddd/domain/location"
	"go-ddd/domain/voyage"
)

// HandlingActivity 表示货物的处理方式和地点，并可用于表达对货物未来可能发生的情况的预测。
type HandlingActivity struct {
	eventType handling.EventType
	location  *location.Location
	voyage    *voyage.Voyage
}

func NewHandlingActivity(eventType handling.EventType, location *location.Location, voyage *voyage.Voyage) (*HandlingActivity, error) {
	if location == nil {
		return nil, errors.New("location is required")
	}

	return &HandlingActivity{
		eventType: eventType,
		location:  location,
		voyage:    voyage,
	}, nil
}

func (h *HandlingActivity) EventType() handling.EventType {
	return h.eventType
}

func (h *HandlingActivity) Location() *location.Location {
	return h.location
}

func (h *HandlingActivity) Voyage() *voyage.Voyage {
	return h.voyage
}
