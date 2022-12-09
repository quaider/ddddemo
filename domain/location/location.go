package location

type Location struct {
	name, unlocode string
}

func NewLocation(name string, unlocode string) *Location {
	return &Location{name: name, unlocode: unlocode}
}

// Unlocode country code and location code concatenated, always upper case.
func (l *Location) Unlocode() string {
	return l.unlocode
}

func (l *Location) Name() string {
	return l.name
}

func (l *Location) SameIdentityAs(that *Location) bool {
	return that != nil && l.unlocode == that.unlocode
}
