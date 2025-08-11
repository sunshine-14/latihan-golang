package controller

import (
	"encoding/json"
	"net/http"
	"trial-go/models"
)

func ListUser(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetAllUsers()
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": err.Error(),
			"status":  "error",
			"data":    nil,
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Users fetched successfully",
		"status":  "success",
		"data":    users,
	})
}
