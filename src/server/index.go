package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mviniciusgc/onfly/src/controllers"
	"github.com/mviniciusgc/onfly/src/db"
	"github.com/mviniciusgc/onfly/src/repositorie"
	"github.com/mviniciusgc/onfly/src/router"
	"github.com/mviniciusgc/onfly/src/service"
	"github.com/spf13/viper"
)

func InitializeServer() {
	// Carregar configuração do arquivo .env
	viper.SetConfigFile(".env")

	// Ler as configurações do .env
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("failed to initialize, error to read .env: %v", err)
	}

	// Inicializa a conexão com o banco de dados
	d := db.InitializeClientBD()

	// Registra as tabelas no banco de dados
	err = d.InitializeTables()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to initialize data base: %w", err))
	}

	// Inicializa o repositório com a dependência de db.IClientDB
	clientRepository := repositorie.InitializeClienteRepository(d)

	// Inicializa o serviço com o repositório
	travelService := service.InitializeService(clientRepository)

	// Inicializa o controller com o serviço
	travelController := controllers.InitializeController(travelService)

	// Registra as rotas no servidor
	router.SetupRoutes(travelController)

	// Lê a porta configurada no .env
	port := viper.GetString("API_PORT")

	// Inicia o servidor HTTP
	fmt.Println("Server is running on port", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Error initialize the server: ", err)
	}
}
