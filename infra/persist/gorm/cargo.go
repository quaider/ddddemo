package gorm

import (
	"context"
	"go-ddd/domain/cargo"
)

// 仅用于确保 CargoRepository 实现了 cargo.Repository 接口
var _ cargo.Repository = (*CargoRepositoryGorm)(nil)

type CargoRepositoryGorm struct {
	cargos map[string]*cargo.Cargo
}

func NewCargoRepositoryGorm() *CargoRepositoryGorm {
	return &CargoRepositoryGorm{
		cargos: make(map[string]*cargo.Cargo, 8),
	}
}

func (c *CargoRepositoryGorm) Find(ctx context.Context, trackingId *cargo.TrackingId) *cargo.Cargo {
	return nil
}

func (c *CargoRepositoryGorm) Save(ctx context.Context, cargo *cargo.Cargo) {

}
