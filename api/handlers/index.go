package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/bitebait/curry/cache"
)

func Index(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	result, err := cache.GetCache()
	if err != nil {
		status = http.StatusNotFound
	}

	json.NewEncoder(w).Encode(result)
}
