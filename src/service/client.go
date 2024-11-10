package service

import repositorie "github.com/mviniciusgc/onfly/src/repositorie"

type TravelService struct {
	clientrepository repositorie.IRepository
}

func InitializeService(r repositorie.IRepository) *TravelService {
	return &TravelService{
		clientrepository: r,
	}
}
