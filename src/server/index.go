package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mviniciusgc/onfly/src/db"
	"github.com/mviniciusgc/onfly/src/routes"
	"github.com/spf13/viper"
)

func InitializeServer() {
	viper.SetConfigFile(".env")
	r := mux.NewRouter()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("failed to initialize, error to read .env: %v", err)
	}

	d := db.InitializeClientBD()

	// Registra a rota /travel
	routes.RegisterAllRoutes(r)

	err = d.InitializeTables()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to initialize data base: %w", err))
	}

	port := viper.GetString("API_PORT")

	fmt.Println("Server is running on port", port)
	err = http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal("Error initialize the server: ", err)
	}

}
