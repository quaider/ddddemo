package domain

// TrackingId 货物跟踪id，最开始时是放在 domain.cargo聚合目录下的
// 编码过程中发现，其他地方也会引用它，可能出现 import cycle的情况，因此权衡考虑
// 决定将其放入通用的 domain 包下，表示 任何其他聚合 对 cargo 通过 id 关联的场景
type TrackingId struct {
	id string
}

func NewTrackingId(id string) *TrackingId {
	return &TrackingId{id: id}
}

func (t *TrackingId) Id() string {
	return t.id
}
