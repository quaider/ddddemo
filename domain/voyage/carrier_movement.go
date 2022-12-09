package voyage

import (
	"errors"
	"go-ddd/domain/location"
	"time"
)

// CarrierMovement 表示从一个地点到另一个地点的航行
type CarrierMovement struct {
	departureLocation *location.Location // 出发地
	arrivalLocation   *location.Location // 到达地
	departureTime     time.Time          // 出发时间
	arrivalTime       time.Time          // 到达时间
}

func NewCarrierMovement(departureLocation, arrivalLocation *location.Location, departureTime, arrivalTime time.Time) (*CarrierMovement, error) {
	if departureLocation == nil || arrivalLocation == nil {
		return nil, errors.New("location can not be null")
	}

	return &CarrierMovement{
		departureLocation: departureLocation,
		arrivalLocation:   arrivalLocation,
		departureTime:     departureTime,
		arrivalTime:       arrivalTime,
	}, nil
}

func (c *CarrierMovement) DepartureLocation() *location.Location {
	return c.departureLocation
}

func (c *CarrierMovement) ArrivalLocation() *location.Location {
	return c.arrivalLocation
}

func (c *CarrierMovement) DepartureTime() time.Time {
	return c.departureTime
}

func (c *CarrierMovement) ArrivalTime() time.Time {
	return c.arrivalTime
}
