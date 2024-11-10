package travel

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mviniciusgc/onfly/src/models"
)

func (c *ServiceController) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	idRequest, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var travelModel models.TravelModel
	err = json.NewDecoder(r.Body).Decode(&travelModel)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	travel, err := c.service.UpdateStatus(idRequest, travelModel.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(travel)
}
