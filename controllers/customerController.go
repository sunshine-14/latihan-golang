package controller

import (
	"encoding/json"
	"net/http"
	"trial-go/models"
)

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := models.GetAllCustomers()
	if err != nil {
		http.Error(w, "Gagal mengambil data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(customers); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
