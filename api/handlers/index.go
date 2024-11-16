package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/bitebait/curry/cache"
)

func Index(w http.ResponseWriter, r *http.Request) {
	const contentType = "application/json"
	setResponseHeaders(w, contentType)

	result, err := cache.GetCache()
	if err != nil {
		http.Error(w, "Cache not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func setResponseHeaders(w http.ResponseWriter, contentType string) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)
}
