package cargo

type TrackingId struct {
	id string
}

func NewTrackingId(id string) *TrackingId {
	return &TrackingId{id: id}
}

func (t *TrackingId) Id() string {
	return t.id
}
