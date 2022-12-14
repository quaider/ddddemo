package mem

import (
	"context"
	"go-ddd/domain"
	"go-ddd/domain/cargo"
	"math/rand"
	"strconv"
)

// 仅用于确保 CargoRepository 实现了 cargo.Repository 接口
var _ cargo.Repository = (*CargoRepository)(nil)

type CargoRepository struct {
	cargos map[string]*cargo.Cargo
}

func NewCargoRepository() *CargoRepository {
	cr := &CargoRepository{
		cargos: make(map[string]*cargo.Cargo),
	}

	return cr
}

func (c *CargoRepository) GenNextId() *domain.TrackingId {
	return domain.NewTrackingId(strconv.Itoa(rand.Intn(1e9)))
}

func (c *CargoRepository) Find(ctx context.Context, trackingId *domain.TrackingId) *cargo.Cargo {
	return nil
}

func (c *CargoRepository) Save(ctx context.Context, cargo *cargo.Cargo) {

}
