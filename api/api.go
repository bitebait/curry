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

		asciiName := config.GetConfig.App.AsciiName
		dbName := config.GetConfig.DB.Name
		apiPort := config.GetConfig.Api.Port
		apiEndpoint := config.GetConfig.Api.Endpoint

		startServer(dbName, asciiName, apiPort, apiEndpoint)
	}()
}

func startServer(dbName, asciiName, apiPort, apiEndpoint string) {
	db.Init(dbName)
	routes.Init()
	fmt.Println(asciiName)
	log.Printf("Running and Listening on :%s\n", apiPort)
	log.Printf("API Endpoint: %s\n", apiEndpoint)
	log.Fatal(http.ListenAndServe(":"+apiPort, nil))
}
