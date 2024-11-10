package repositorie

import "github.com/mviniciusgc/onfly/src/entity"

type IRepository interface {
	Create(travel entity.TravelEntity) (id *int64, err error)
	ListAll(status string) ([]entity.TravelEntity, error)
	ListById(id int) (*entity.TravelEntity, error)
	UpdateStatus(id int, status string) (*entity.TravelEntity, error)
}
