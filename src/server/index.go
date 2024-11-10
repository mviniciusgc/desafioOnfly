package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	controllers "github.com/mviniciusgc/onfly/src/controllers/travel"
	"github.com/mviniciusgc/onfly/src/db"
	repositorieLog "github.com/mviniciusgc/onfly/src/repositorie/logs"
	repositorie "github.com/mviniciusgc/onfly/src/repositorie/travel"
	"github.com/mviniciusgc/onfly/src/router"
	service "github.com/mviniciusgc/onfly/src/service/travel"
	"github.com/spf13/viper"
)

func InitializeServer() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed to initialize, error to read .env: %v", err)
	}

	database := db.InitializeClientBD()
	if err := database.InitializeTables(); err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	clientRepository := repositorie.InitializeClienteRepository(database)
	clientLogRepository := repositorieLog.InitializeClienteRepository(database)
	travelService := service.InitializeService(clientRepository, clientLogRepository)
	travelController := controllers.InitializeController(travelService)

	r := mux.NewRouter()
	router.SetupRoutes(r, travelController)

	port := viper.GetString("API_PORT")
	fmt.Printf("Server is running on port %s\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("Error initializing the server: %v", err)
	}
}
