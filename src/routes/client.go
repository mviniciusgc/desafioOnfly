package routes

import "github.com/gorilla/mux"

// Função para registrar todas as rotas do projeto
func RegisterAllRoutes(r *mux.Router) {
	// Instancia as rotas específicas
	travelRoute := &TravelRoute{}

	// Adiciona cada rota ao roteador principal
	travelRoute.RegisterRoutes(r)
}
