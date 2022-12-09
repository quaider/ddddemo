package handling

// EventType 表示 是否 需要或不需要 a carrier movement
type EventType int

const (
	LOAD EventType = iota
	UNLOAD
	RECEIVE
	CLAIM
	CUSTOMS
)

func (e EventType) BoolValue() bool {
	return e < 2
}
