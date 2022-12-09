package service

import (
	"go-ddd/domain/cargo"
	"go-ddd/domain/location"
	"go-ddd/domain/voyage"
	"go-ddd/infra/persist/gorm"
	"go-ddd/infra/persist/mem"
	"time"
)

// CargoConfig 表示 CargoService 的配置
type CargoConfig func(cs *CargoService) error

// CargoService 应用服务
type CargoService struct {
	cargos  cargo.Repository
	voyages voyage.Repository
}

func (s *CargoService) Cargos() cargo.Repository {
	return s.cargos
}

func (s *CargoService) Voyages() voyage.Repository {
	return s.voyages
}

// NewCargoService 根据 配置序列 创建 CargoService 实例
func NewCargoService(configs ...CargoConfig) (*CargoService, error) {

	cs := &CargoService{}
	// 应用服务的所有配置
	for _, cfg := range configs {
		err := cfg(cs)
		if err != nil {
			return nil, err
		}
	}

	return cs, nil
}

func WithCargoRepository(cr cargo.Repository) CargoConfig {

	return func(cs *CargoService) error {
		cs.cargos = cr
		return nil
	}
}

func WithVoyagesRepository(vr voyage.Repository) CargoConfig {

	return func(cs *CargoService) error {
		cs.voyages = vr
		return nil
	}
}

// WithMemoryCargoRepository cargo 内存仓储配置
func WithMemoryCargoRepository() CargoConfig {
	cr := mem.NewCargoRepository()
	return WithCargoRepository(cr)
}

// WithGormCargoRepository cargo gorm仓储配置
func WithGormCargoRepository() CargoConfig {
	cr := gorm.NewCargoRepositoryGorm()
	return WithCargoRepository(cr)
}

// WithMemoryVoyageRepository voyage 内存仓储配置
func WithMemoryVoyageRepository() CargoConfig {
	cr := mem.NewCargoRepository()
	return WithCargoRepository(cr)
}

// WithGormVoyageRepository voyage gorm仓储配置
func WithGormVoyageRepository() CargoConfig {
	cr := gorm.NewCargoRepositoryGorm()
	return WithCargoRepository(cr)
}

func (s *CargoService) CreateCargo(origin, destination *location.Location, arrivalDeadline time.Time) (*cargo.Cargo, error) {
	specification, err := cargo.NewRouteSpecification(origin, destination, arrivalDeadline)
	if err != nil {
		return nil, err
	}

	trackingId := s.cargos.GenNextId()

	return cargo.NewCargo(trackingId, specification)
}
