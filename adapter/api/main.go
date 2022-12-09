package main

import (
	"github.com/davecgh/go-spew/spew"
	"go-ddd/adapter/service"
	domainCargo "go-ddd/domain/cargo"
	"go-ddd/domain/handling"
	"go-ddd/domain/location"
	"go-ddd/domain/voyage"
	"time"
)

func main() {
	cargoService, err := service.NewCargoService(
		service.WithGormCargoRepository(),
		service.WithMemoryVoyageRepository(),
	)

	if err != nil {
		panic(err)
	}

	shLocation := location.NewLocation("Shanghai", "021")
	whLocation := location.NewLocation("Wuhan", "027")
	hsLocation := location.NewLocation("Huangshi", "0714")

	cargo, err := cargoService.CreateCargo(whLocation, hsLocation)

	if err != nil {
		panic(err)
	}

	goodI := domainCargo.NewItineraryEmpty()
	err = cargo.AssignToRoute(goodI)
	if err != nil {
		panic(err)
	}

	evt, err := handling.NewEventWithoutVoyage(handling.RECEIVE, shLocation, time.Now().Add(240*time.Hour), time.Now(), cargo.TrackingId().Id())
	if err != nil {
		panic(err)
	}

	events := make([]*handling.Event, 0)
	events = append(events, evt)
	if err = cargo.DeriveDeliveryProgress(handling.NewHistory(events)); err != nil {
		panic(err)
	}

	voy := voyage.NewBuilder("0123", shLocation).
		AddMovement(whLocation, time.Now(), time.Now()).
		Build()

	evt, err = handling.NewEvent(handling.UNLOAD, voy, shLocation, time.Now().Add(480*time.Hour), time.Now(), cargo.TrackingId().Id())
	if err != nil {
		panic(err)
	}

	events = append(events, evt)
	if err = cargo.DeriveDeliveryProgress(handling.NewHistory(events)); err != nil {
		panic(err)
	}

	spew.Config = spew.ConfigState{
		Indent:                  "\t",
		DisablePointerAddresses: true,
	}
	spew.Dump(cargo)
}
