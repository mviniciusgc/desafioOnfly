package repositorie

import "github.com/mviniciusgc/onfly/src/entity"

type IRepository interface {
	Create(travel entity.TravelEntity) (id *int64, err error)
}
