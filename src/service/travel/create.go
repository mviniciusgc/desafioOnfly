package travel

import (
	"errors"
	"strconv"

	"github.com/mviniciusgc/onfly/src/entity"
	"github.com/mviniciusgc/onfly/src/enum"
	"github.com/mviniciusgc/onfly/src/models"
)

func (s *TravelService) Create(travel *models.TravelModel) (id *int64, err error) {

	if travel.RequesterName == "" || travel.Destination == "" {
		s.logepository.CreateLog(enum.Error, "error create travel")
		return nil, errors.New("applicant name and destination are required")
	}
	travelEntity := ConvertToTravelEntity(*travel)

	response, err := s.clientrepository.Create(travelEntity)
	if err != nil {
		messageLog := "error create travel: " + err.Error()
		s.logepository.CreateLog(enum.Error, messageLog)
		return nil, err
	}

	str := ""
	if response != nil {
		str = strconv.FormatInt(*response, 10)
	}

	messageLog := "success in recording the trip request ID: " + str
	s.logepository.CreateLog(enum.Info, messageLog)

	return response, nil
}

func ConvertToTravelEntity(t models.TravelModel) entity.TravelEntity {
	return entity.TravelEntity{
		RequesterName: t.RequesterName,
		Destination:   t.Destination,
		DepartureDate: t.DepartureDate,
		ReturnDate:    t.ReturnDate,
		Status:        t.Status,
		CreatedAt:     t.CreatedAt,
	}
}
