package travel

import (
	"strconv"

	"github.com/mviniciusgc/onfly/src/enum"
	"github.com/mviniciusgc/onfly/src/models"
)

func (s *TravelService) UpdateStatus(id int, status string) (*models.TravelModel, error) {

	travelEntitie, err := s.clientrepository.UpdateStatus(id, status)
	if err != nil {
		messageLog := "error to update travel: " + err.Error()
		s.logepository.CreateLog(enum.Info, messageLog)
		return nil, err
	}

	travelModel := convertToModel(*travelEntitie)

	str := strconv.Itoa(id)

	messageLog := "success in recording the trip request ID: " + str + "status: " + status
	s.logepository.CreateLog(enum.Info, messageLog)

	return &travelModel, nil
}
