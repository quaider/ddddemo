package voyage

import (
	"go-ddd/domain/location"
	"time"
)

type Voyage struct {
	voyageNumber string
	schedule     *Schedule
}

func NewVoyage(voyageNumber string, schedule *Schedule) *Voyage {
	return &Voyage{voyageNumber: voyageNumber, schedule: schedule}
}

func (v *Voyage) VoyageNumber() string {
	return v.voyageNumber
}

func (v *Voyage) Schedule() *Schedule {
	return v.schedule
}

func (v *Voyage) SameIdentityAs(that *Voyage) bool {
	return that != nil && v.voyageNumber == that.voyageNumber
}

// Builder 演示 go中如何实现 Builder 模式
type Builder struct {
	carrierMovements  []*CarrierMovement
	voyageNumber      string
	departureLocation *location.Location
}

func NewBuilder(voyageNumber string, departureLocation *location.Location) *Builder {
	return &Builder{voyageNumber: voyageNumber, departureLocation: departureLocation}
}

func (b *Builder) AddMovement(arrivalLocation *location.Location, departureTime, arrivalTime time.Time) *Builder {
	movement, _ := NewCarrierMovement(b.departureLocation, arrivalLocation, departureTime, arrivalTime)
	b.carrierMovements = append(b.carrierMovements, movement)

	// Next departure location is the same as this arrival location
	b.departureLocation = arrivalLocation

	return b
}

func (b *Builder) Build() *Voyage {
	schedule, _ := newSchedule(b.carrierMovements)
	return NewVoyage(b.voyageNumber, schedule)
}
