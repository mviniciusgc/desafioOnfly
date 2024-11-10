package service

import (
	"errors"

	"github.com/mviniciusgc/onfly/src/entity"
	"github.com/mviniciusgc/onfly/src/models"
)

func (s *TravelService) Create(travel *models.TravelModel) (id *int64, err error) {
	// fmt.Println("service")
	if travel.RequesterName == "" || travel.Destination == "" {
		return nil, errors.New("nome do solicitante e destino são obrigatórios")
	}
	travel.Status = "solicitado"

	travelEntity := ConvertToTravelEntity(*travel)
	// fmt.Println(travel)
	return s.clientrepository.Create(travelEntity)
}

// Função para converter de Travel (model) para TravelEntity (entity)
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
