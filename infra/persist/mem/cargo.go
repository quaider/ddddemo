package mem

import (
	"context"
	"go-ddd/domain/cargo"
)

// 仅用于确保 CargoRepository 实现了 cargo.Repository 接口
var _ cargo.Repository = (*CargoRepository)(nil)

type CargoRepository struct {
	cargos map[string]*cargo.Cargo
}

func NewCargoRepository() *CargoRepository {
	cr := &CargoRepository{
		cargos: make(map[string]*cargo.Cargo, 8),
	}

	return cr
}

func (c *CargoRepository) Find(ctx context.Context, trackingId *cargo.TrackingId) *cargo.Cargo {
	return nil
}

func (c *CargoRepository) Save(ctx context.Context, cargo *cargo.Cargo) {

}
