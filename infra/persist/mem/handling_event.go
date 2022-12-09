package mem

import (
	"go-ddd/domain"
	"go-ddd/domain/handling"
)

var _ handling.Repository = (*HandlingEventRepository)(nil)

type HandlingEventRepository struct {
	events map[string]*handling.Event
}

func NewHandlingEventRepository() *HandlingEventRepository {
	cr := &HandlingEventRepository{
		events: make(map[string]*handling.Event),
	}

	return cr
}

func (v *HandlingEventRepository) Save(event *handling.Event) error {
	return nil
}

func (v *HandlingEventRepository) LookupHandlingHistoryOfCargo(trackingId *domain.TrackingId) *handling.History {
	return nil
}
