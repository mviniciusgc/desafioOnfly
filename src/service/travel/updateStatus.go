package travel

import (
	"github.com/mviniciusgc/onfly/src/models"
)

func (s *TravelService) UpdateStatus(id int, status string) (*models.TravelModel, error) {

	travelEntitie, err := s.clientrepository.UpdateStatus(id, status)
	if err != nil {
		return nil, err
	}

	travelModel := convertToModel(*travelEntitie)

	return &travelModel, nil
}
