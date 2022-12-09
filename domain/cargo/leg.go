package cargo

import (
	"go-ddd/domain/location"
	"go-ddd/domain/voyage"
	"time"
)

type Leg struct {
	voyage         *voyage.Voyage
	loadLocation   *location.Location
	unloadLocation *location.Location
	loadTime       time.Time
	unloadTime     time.Time
}

func NewLeg(voyage *voyage.Voyage, loadLocation *location.Location, unloadLocation *location.Location, loadTime time.Time, unloadTime time.Time) *Leg {
	return &Leg{voyage: voyage, loadLocation: loadLocation, unloadLocation: unloadLocation, loadTime: loadTime, unloadTime: unloadTime}
}

func (l *Leg) Voyage() *voyage.Voyage {
	return l.voyage
}

func (l *Leg) LoadLocation() *location.Location {
	return l.loadLocation
}

func (l *Leg) UnloadLocation() *location.Location {
	return l.unloadLocation
}

func (l *Leg) LoadTime() time.Time {
	return l.loadTime
}

func (l *Leg) UnloadTime() time.Time {
	return l.unloadTime
}
