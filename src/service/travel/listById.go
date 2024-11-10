package travel

import (
	"github.com/mviniciusgc/onfly/src/enum"
	"github.com/mviniciusgc/onfly/src/models"
)

func (s *TravelService) ListById(id int) (*models.TravelModel, error) {

	travelEntitie, err := s.clientrepository.ListById(id)
	if err != nil {

		messageLog := "error to list by id travel: " + err.Error()
		s.logepository.CreateLog(enum.Error, messageLog)

		return nil, err
	}

	travelModel := convertToModel(*travelEntitie)

	messageLog := "success in list by id"
	s.logepository.CreateLog(enum.Info, messageLog)

	return &travelModel, nil
}
