package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/bitebait/curry/api/db"
	"github.com/bitebait/curry/api/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	cache := &models.Cache{}
	status := http.StatusOK

	result := db.Database.Preload("Stores").Last(&cache)
	if result.RowsAffected <= 0 {
		status = http.StatusNotFound
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(cache)
}
