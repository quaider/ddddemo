package cargo

import (
	"context"
	"go-ddd/domain"
)

// Repository 表示 Cargo聚合的仓储
type Repository interface {
	GenNextId() *domain.TrackingId

	Find(ctx context.Context, trackingId *domain.TrackingId) *Cargo

	Save(ctx context.Context, cargo *Cargo)
}
