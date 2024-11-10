package travel

import (
	logerepositorie "github.com/mviniciusgc/onfly/src/repositorie/logs"
	repositorie "github.com/mviniciusgc/onfly/src/repositorie/travel"
)

type TravelService struct {
	clientrepository repositorie.IRepository
	logepository     logerepositorie.IRepository
}

func InitializeService(r repositorie.IRepository, lr logerepositorie.IRepository) *TravelService {
	return &TravelService{
		clientrepository: r,
		logepository:     lr,
	}
}
