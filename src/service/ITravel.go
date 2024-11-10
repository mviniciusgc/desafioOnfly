package service

import (
	"github.com/mviniciusgc/onfly/src/models"
)

type IService interface {
	Create(travel *models.TravelModel) (id *int64, err error)
}
