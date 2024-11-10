package travel

import repositorie "github.com/mviniciusgc/onfly/src/repositorie/travel"

type TravelService struct {
	clientrepository repositorie.IRepository
}

func InitializeService(r repositorie.IRepository) *TravelService {
	return &TravelService{
		clientrepository: r,
	}
}
