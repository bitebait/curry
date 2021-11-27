package handlers

import (
	"curry/api/db"
	"curry/api/models"
	"encoding/json"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	dB := db.GetDB()

	cache := &models.Cache{}
	status := http.StatusOK

	result := dB.Preload("Stores").Last(&cache)
	if result.RowsAffected <= 0 {
		status = http.StatusNotFound
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(cache)
}
