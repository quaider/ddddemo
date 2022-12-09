package location

type Repository interface {
	FindLocation(unlocode string) *Location
}
