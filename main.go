package main

import (
	"github.com/mviniciusgc/onfly/src/server"
)

func main() {

	server.InitializeServer()

	// r := mux.NewRouter()

	// // Registra a rota /travel
	// routes.RegisterAllRoutes(r)

	// // Inicia o servidor
	// log.Println("Servidor rodando na porta 8080")
	// log.Fatal(http.ListenAndServe(":8080", r))
}
