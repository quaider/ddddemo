package consumer

import (
	"context"
	"fmt"
	"github.com/asaskevich/EventBus"
	"github.com/davecgh/go-spew/spew"
	"github.com/golobby/container/v3"
	"go-ddd/domain"
	"go-ddd/domain/handling"
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

func StartLocal() {
	var bus EventBus.Bus
	container.MustResolve(container.Global, &bus)
	err := bus.Subscribe("cargo:cargoHandled", func(event *handling.Event) {
		fmt.Println("receive cargo:cargoHandled event")
		spew.Dump(event)
	})
	if err != nil {
		panic(err)
	}
}
