package router

import (
	"net/http"

	"github.com/mviniciusgc/onfly/src/controllers"
)

func SetupRoutes(controller *controllers.ServiceController) {
	http.HandleFunc("/travel", controller.Create)
}
