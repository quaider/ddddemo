package cargo

type RoutingStatus int

const (
	NotRouted RoutingStatus = iota
	Routed
	MisRouted
)

type Test struct {
	status RoutingStatus
}
