package travel

import (
	"github.com/mviniciusgc/onfly/src/models"
)

type IService interface {
	Create(travel *models.TravelModel) (id *int64, err error)
	ListAll(status string) ([]models.TravelModel, error)
	ListById(id int) (*models.TravelModel, error)
	UpdateStatus(id int, status string) (*models.TravelModel, error)
}
