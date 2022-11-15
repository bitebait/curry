package api

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/bitebait/curry/api/db"
	"github.com/bitebait/curry/api/routes"
	"github.com/bitebait/curry/config"
)

func Init(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()

		db.Init(config.GetConfig.DB.Name)

		routes.Init()

		fmt.Println(config.GetConfig.App.AsciiName)
		log.Printf("Running and Listening on :%s\n", config.GetConfig.Api.Port)
		log.Printf("API Endpoint: %s\n", config.GetConfig.Api.Endpoint)
		log.Fatal(http.ListenAndServe(":"+config.GetConfig.Api.Port, nil))
	}()

}
