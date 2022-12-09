package handling

// EventType 表示 是否 需要或不需要 a carrier movement
type EventType bool

const (
	LOAD    EventType = true
	UNLOAD            = true
	RECEIVE           = false
	CLAIM             = false
	CUSTOMS           = false
)
