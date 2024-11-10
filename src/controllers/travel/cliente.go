package travel

import service "github.com/mviniciusgc/onfly/src/service/travel"

type ServiceController struct {
	service service.IService
}

func InitializeController(s service.IService) *ServiceController {

	return &ServiceController{
		service: s,
	}
}
