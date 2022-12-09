package voyage

type Repository interface {
	FindVoyage(voyageNumber string) *Voyage
}
