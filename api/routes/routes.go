package routes

import (
	"net/http"

	"github.com/bitebait/curry/api/handlers"
	"github.com/bitebait/curry/config"
)

func Init() {
	cfg := config.GetConfig()
	http.HandleFunc(cfg.Api.Endpoint, handlers.Index)
}
