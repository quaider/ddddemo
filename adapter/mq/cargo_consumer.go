package mq

import (
	"context"
	"go-ddd/domain"
	"go-ddd/infra/persist/mem"
)

// onMessage mock mq consumer
func onMessage(trackingId string) {
	cargoRepository := mem.NewCargoRepository()
	handlingEventRepository := mem.NewHandlingEventRepository()
	newTrackingId := domain.NewTrackingId(trackingId)
	cargo := cargoRepository.Find(context.Background(), newTrackingId)
	if cargo == nil {
		return
	}

	handlingHistory := handlingEventRepository.LookupHandlingHistoryOfCargo(newTrackingId)
	cargo.DeriveDeliveryProgress(handlingHistory)

	if cargo.Delivery().Misdirected() {
		// publish cargo was mis direct event
	}

	if cargo.Delivery().IsUnloadedAtDestination() {
		// publish cargo has arrived event
	}

	cargoRepository.Save(context.Background(), cargo)
}
