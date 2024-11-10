// package server

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// 	controllers "github.com/mviniciusgc/onfly/src/controllers/travel"
// 	"github.com/mviniciusgc/onfly/src/db"
// 	repositorie "github.com/mviniciusgc/onfly/src/repositorie/travel"
// 	"github.com/mviniciusgc/onfly/src/router"
// 	service "github.com/mviniciusgc/onfly/src/service/travel"
// 	"github.com/spf13/viper"
// )

// func InitializeServer() {
// 	// Carregar configuração do arquivo .env
// 	viper.SetConfigFile(".env")
// 	r := mux.NewRouter()

// 	// Ler as configurações do .env
// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		log.Fatalf("failed to initialize, error to read .env: %v", err)
// 	}

// 	// Inicializa a conexão com o banco de dados
// 	d := db.InitializeClientBD()

// 	// Registra as tabelas no banco de dados
// 	err = d.InitializeTables()
// 	if err != nil {
// 		log.Fatal(fmt.Errorf("failed to initialize data base: %w", err))
// 	}

// 	// Inicializa o repositório com a dependência de db.IClientDB
// 	clientRepository := repositorie.InitializeClienteRepository(d)

// 	// Inicializa o serviço com o repositório
// 	travelService := service.InitializeService(clientRepository)

// 	// Inicializa o controller com o serviço
// 	travelController := controllers.InitializeController(travelService)

// 	// Registra as rotas no servidor
// 	router.SetupRoutes(r, travelController)

// 	// Lê a porta configurada no .env
// 	port := viper.GetString("API_PORT")

// 	// Inicia o servidor HTTP
// 	fmt.Println("Server is running on port", port)
// 	err = http.ListenAndServe(port, nil)
// 	if err != nil {
// 		log.Fatal("Error initialize the server: ", err)
// 	}
// }

package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	controllers "github.com/mviniciusgc/onfly/src/controllers/travel"
	"github.com/mviniciusgc/onfly/src/db"
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
	travelService := service.InitializeService(clientRepository)
	travelController := controllers.InitializeController(travelService)

	r := mux.NewRouter()
	router.SetupRoutes(r, travelController)

	port := viper.GetString("API_PORT")
	fmt.Printf("Server is running on port %s\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("Error initializing the server: %v", err)
	}
}
