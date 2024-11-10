package travel

import (
	"errors"

	"github.com/mviniciusgc/onfly/src/entity"
	"github.com/mviniciusgc/onfly/src/models"
)

func convertToModel(travelEntity entity.TravelEntity) models.TravelModel {
	id := int64(travelEntity.ID)
	return models.TravelModel{
		ID:            &id,
		RequesterName: travelEntity.RequesterName,
		Destination:   travelEntity.Destination,
		DepartureDate: travelEntity.DepartureDate,
		ReturnDate:    travelEntity.ReturnDate,
		Status:        travelEntity.Status,
		CreatedAt:     travelEntity.CreatedAt,
	}
}

func (s *TravelService) ListAll(status string) ([]models.TravelModel, error) {

	travelEntities, err := s.clientrepository.ListAll(status)
	if err != nil {
		return nil, err
	}

	if len(travelEntities) == 0 {
		return nil, errors.New("no orders found")
	}

	var travelModels []models.TravelModel
	for _, entity := range travelEntities {
		travelModels = append(travelModels, convertToModel(entity))
	}

	return travelModels, nil
}
