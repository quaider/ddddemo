package main

import (
	"github.com/davecgh/go-spew/spew"
	"go-ddd/adapter/service"
	domainCargo "go-ddd/domain/cargo"
	"go-ddd/domain/handling"
	"go-ddd/domain/location"
	"go-ddd/domain/voyage"
	"go-ddd/infra/persist/mem"
	"time"
)

var (
	shLocation       = location.NewLocation("上海", "200000")
	kunShanLocation  = location.NewLocation("昆山", "200001")
	zjLocation       = location.NewLocation("浙江", "310000")
	whLocation       = location.NewLocation("武汉", "430000")
	bjLocation       = location.NewLocation("北京", "100000")
	emptyItineraries = make([]*domainCargo.Itinerary, 0)

	cargoService, _ = service.NewCargoService(
		service.WithGormCargoRepository(),
		service.WithMemoryVoyageRepository(),
	)

	handlingEventService = service.NewHandlingEventServiceImpl(
		mem.NewLocationRepository(), mem.NewVoyageRepository(), mem.NewHandlingEventRepository())

	voya_100 = voyage.NewBuilder("V100", shLocation).
			AddMovement(kunShanLocation, toDate("2022-12-09"), toDate("2022-12-11")).
			AddMovement(zjLocation, toDate("2022-12-12"), toDate("2022-12-15")).
			Build()

	voya_200 = voyage.NewBuilder("V200", shLocation).
			AddMovement(zjLocation, toDate("2022-12-16"), toDate("2022-12-18")).
			AddMovement(whLocation, toDate("2022-12-19"), toDate("2022-12-21")).
			AddMovement(bjLocation, toDate("2022-12-22"), toDate("2022-12-25")).
			Build()
)

func toDate(dateStr string) time.Time {
	parse, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		panic(err)
	}

	return parse
}

func bookNewCargo(cargoService *service.CargoService) *domainCargo.Cargo {
	cargo, err := cargoService.CreateCargo(shLocation, bjLocation, toDate("2022-12-26"))
	if err != nil {
		panic(err)
	}

	return cargo
}

// requestItineraryFromMockService 请求航线
func requestItineraryFromMockService(cargo *domainCargo.Cargo) []*domainCargo.Itinerary {
	if cargo == nil {
		return emptyItineraries
	}

	legs := make([]*domainCargo.Leg, 0)

	// 从上海出发的路线
	if cargo.RouteSpecification().Origin().SameIdentityAs(shLocation) {
		leg1 := domainCargo.NewLeg(voya_100, shLocation, zjLocation, toDate("2022-12-09"), toDate("2022-12-15"))
		leg2 := domainCargo.NewLeg(voya_200, zjLocation, whLocation, toDate("2022-12-16"), toDate("2022-12-21"))
		leg3 := domainCargo.NewLeg(voya_100, whLocation, bjLocation, toDate("2022-12-22"), toDate("2022-12-25"))

		legs = append(legs, leg1, leg2, leg3)
	}

	itinerary, err := domainCargo.NewItinerary(legs)
	if err != nil {
		panic(err)
	}

	itineraries := make([]*domainCargo.Itinerary, 0, 1)
	itineraries = append(itineraries, itinerary)

	return itineraries
}

func selectPreferedItinerary(itineraries []*domainCargo.Itinerary) *domainCargo.Itinerary {
	return itineraries[0]
}

func init() {
	spew.Config = spew.ConfigState{
		Indent:                  "\t",
		DisablePointerAddresses: true,
	}
}

func main() {

	// 1.0 预定货运， 从上海运到北京
	cargo := bookNewCargo(cargoService)

	// 1.1 查询满足路径规格的所有航线
	itineraries := requestItineraryFromMockService(cargo)

	// 1.2 选定一个合适的航线
	itinerary := selectPreferedItinerary(itineraries)

	// 1.3 分配选定的航线
	err := cargo.AssignToRoute(itinerary)
	if err != nil {
		panic(err)
	}

	// 1.4 handling
	_ = handlingEventService.RegisterHandlingEvent(
		toDate("2022-12-09"), *cargo.TrackingId(), "", shLocation.Unlocode(), handling.RECEIVE)

	_ = handlingEventService.RegisterHandlingEvent(
		toDate("2022-12-09"), *cargo.TrackingId(), voya_100.VoyageNumber(), shLocation.Unlocode(), handling.LOAD)

	spew.Dump(cargo)
}
