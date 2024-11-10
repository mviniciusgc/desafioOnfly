package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mviniciusgc/onfly/src/models"
	"github.com/mviniciusgc/onfly/src/service"
)

type ServiceController struct {
	service service.IService
}

func (c *ServiceController) Create(w http.ResponseWriter, r *http.Request) {
	var travel models.TravelModel
	err := json.NewDecoder(r.Body).Decode(&travel)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// fmt.(travel)
	// fmt.Printf("%+v\n", travel)
	//fmt.Printf("Decoded Travel struct: %+v\n", r.Body)

	_, err = c.service.Create(&travel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Travel request created successfully"})
}
