package routes

import (
	"curry/api/handlers"
	"curry/config"
	"net/http"
)

func Init() {
	cfg := config.GetConfig()
	http.HandleFunc(cfg.Api.Endpoint, handlers.Index)
}
