package mem

import (
	"go-ddd/domain/voyage"
)

var _ voyage.Repository = (*VoyageRepository)(nil)

type VoyageRepository struct {
	voyages map[string]*voyage.Voyage
}

func NewVoyageRepository() *VoyageRepository {
	cr := &VoyageRepository{
		voyages: make(map[string]*voyage.Voyage),
	}

	return cr
}

func (v *VoyageRepository) FindVoyage(voyageNumber string) *voyage.Voyage {
	return nil
}
