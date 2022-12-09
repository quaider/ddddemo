package voyage

import "errors"

// Schedule 表示 航行安排，即一系列的 CarrierMovement
type Schedule struct {
	carrierMovements []*CarrierMovement
}

func newSchedule(carrierMovements []*CarrierMovement) (*Schedule, error) {
	if carrierMovements == nil || len(carrierMovements) == 0 {
		return nil, errors.New("carrierMovements can not be empty")
	}

	return &Schedule{
		carrierMovements: carrierMovements,
	}, nil
}
