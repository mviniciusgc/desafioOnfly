package router

import (
	"net/http"

	"github.com/gorilla/mux"
	controllerTravel "github.com/mviniciusgc/onfly/src/controllers/travel"
)

func SetupRoutes(router *mux.Router, controller *controllerTravel.ServiceController) {

	router.HandleFunc("/travel", controller.Create).Methods(http.MethodPost)
	router.HandleFunc("/travel/listAll", controller.ListAll).Methods(http.MethodGet)
	router.HandleFunc("/travel/{id}", controller.ListById).Methods(http.MethodGet)
	router.HandleFunc("/travel/updateStatus/{id}", controller.UpdateStatus).Methods(http.MethodPatch)

}
