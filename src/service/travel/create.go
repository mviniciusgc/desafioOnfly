package travel

import (
	"errors"

	"github.com/mviniciusgc/onfly/src/entity"
	"github.com/mviniciusgc/onfly/src/models"
)

func (s *TravelService) Create(travel *models.TravelModel) (id *int64, err error) {

	if travel.RequesterName == "" || travel.Destination == "" {
		return nil, errors.New("applicant name and destination are required")
	}
	travelEntity := ConvertToTravelEntity(*travel)

	return s.clientrepository.Create(travelEntity)
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
