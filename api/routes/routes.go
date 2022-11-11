package routes

import (
	"net/http"

	"github.com/bitebait/curry/api/handlers"
	"github.com/bitebait/curry/config"
)

func Init() {
	http.HandleFunc(config.GetConfig.Api.Endpoint, handlers.Index)
}
