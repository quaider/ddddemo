package mem

import (
	"go-ddd/domain/location"
)

var _ location.Repository = (*LocationRepository)(nil)

type LocationRepository struct {
	locations map[string]*location.Location
}

func NewLocationRepository() *LocationRepository {
	cr := &LocationRepository{
		locations: make(map[string]*location.Location),
	}

	return cr
}

func (l *LocationRepository) FindLocation(unlocode string) *location.Location {
	return nil
}
