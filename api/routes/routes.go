package routes

import (
	"github.com/bitebait/curry/api/handlers"
	"github.com/bitebait/curry/config"
	"net/http"
)

func Init() {
	cfg := config.GetConfig()
	http.HandleFunc(cfg.Api.Endpoint, handlers.Index)
}
