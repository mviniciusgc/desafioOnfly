package travel

import (
	"github.com/mviniciusgc/onfly/src/models"
)

func (s *TravelService) ListById(id int) (*models.TravelModel, error) {

	travelEntitie, err := s.clientrepository.ListById(id)
	if err != nil {
		return nil, err
	}

	travelModel := convertToModel(*travelEntitie)

	return &travelModel, nil
}
