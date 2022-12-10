package ioc

import (
	"github.com/golobby/container/v3"
	"go-ddd/adapter/service"
	"go-ddd/domain/cargo"
	"go-ddd/domain/handling"
	"go-ddd/domain/location"
	"go-ddd/domain/voyage"
	"go-ddd/infra/persist/mem"
)

func Bootstrap() {
	c := container.Global

	// =========== install configs ===========

	// =========== install repositories ===========
	container.MustSingleton(c, func() cargo.Repository {
		return mem.NewCargoRepository()
	})

	container.MustSingleton(c, func() voyage.Repository {
		return mem.NewVoyageRepository()
	})

	container.MustSingleton(c, func() handling.Repository {
		return mem.NewHandlingEventRepository()
	})

	container.MustSingleton(c, func() location.Repository {
		return mem.NewLocationRepository()
	})

	// =========== install factories ===========
	// 注册时是指针，resolve时也必须是指针
	container.MustSingleton(c, func(lr location.Repository, vr voyage.Repository) *handling.EventFactory {
		return handling.NewEventFactory(vr, lr)
	})

	// =========== install domain services ===========

	// =========== install application services ===========
	container.MustSingleton(c, func(cr cargo.Repository, vr voyage.Repository) *service.CargoService {
		return service.NewCargoService(cr, vr)
	})
}
