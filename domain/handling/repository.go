package handling

type Repository interface {
	Save(event *Event) error
}
