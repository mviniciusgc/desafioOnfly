// routes/travel.go
package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Estrutura para representar o corpo da requisição JSON
type TravelRequest struct {
	Destination string `json:"destination"`
	Duration    int    `json:"duration"`
}

// Rota para a funcionalidade de viagem
type TravelRoute struct{}

// Função que lida com a requisição POST para a rota /travel
func (tr *TravelRoute) handleTravel(w http.ResponseWriter, r *http.Request) {
	var travelReq TravelRequest

	// Decodifica o corpo da requisição JSON na estrutura TravelRequest
	err := json.NewDecoder(r.Body).Decode(&travelReq)
	if err != nil {
		http.Error(w, "Erro ao processar a requisição", http.StatusBadRequest)
		return
	}

	// Lógica para tratar a requisição
	log.Printf("Destino: %s, Duração: %d dias", travelReq.Destination, travelReq.Duration)

	// Responde com um JSON de confirmação
	response := map[string]string{"message": "Requisição de viagem recebida com sucesso"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Implementa RegisterRoutes para TravelRoute
func (tr *TravelRoute) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/travel", tr.handleTravel).Methods("POST")
}
