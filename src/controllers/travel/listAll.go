package travel

import (
	"encoding/json"
	"net/http"
)

func (c *ServiceController) ListAll(w http.ResponseWriter, r *http.Request) {

	status := r.URL.Query().Get("status")

	travels, err := c.service.ListAll(status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(travels)
}
