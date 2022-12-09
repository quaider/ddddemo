package cargo

// TransportStatus 表示货物运送状态
type TransportStatus int

const (
	NOT_RECEIVED = iota
	IN_PORT
	ONBOARD_CARRIER
	CLAIMED
	UNKNOWN
)
