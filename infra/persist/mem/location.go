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
		locations: map[string]*location.Location{
			"200000": location.NewLocation("上海", "200000"),
			"200001": location.NewLocation("昆山", "200001"),
			"310000": location.NewLocation("浙江", "310000"),
			"430000": location.NewLocation("武汉", "430000"),
			"100000": location.NewLocation("北京", "100000"),
		},
	}

	return cr
}

func (l *LocationRepository) FindLocation(unlocode string) *location.Location {
	if lo, ok := l.locations[unlocode]; ok {
		return lo
	}

	return nil
}
