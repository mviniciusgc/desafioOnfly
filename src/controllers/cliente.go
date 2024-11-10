package controllers

import "github.com/mviniciusgc/onfly/src/service"

// Alteração na função para retornar *ServiceController
func InitializeController(s service.IService) *ServiceController {
	// Criando o controlador de serviço (ServiceController)
	return &ServiceController{
		service: s,
	}
}

// package controllers

// import (
// 	"github.com/mviniciusgc/onfly/src/service"
// )

// type TravelController struct {
// 	ClientService service.IService
// }

// func InitializeController(s service.IService) *TravelController {
// 	return &TravelController{
// 		ClientService: s,
// 	}
// }
