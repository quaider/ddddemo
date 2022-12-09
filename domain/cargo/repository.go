package cargo

import "context"

// Repository 表示 Cargo聚合的仓储
type Repository interface {
	Find(ctx context.Context, trackingId *TrackingId) *Cargo

	Save(ctx context.Context, cargo *Cargo)
}
