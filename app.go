package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/bitebait/curry/api/db"
	"github.com/bitebait/curry/api/models"
	"github.com/bitebait/curry/api/routes"
	"github.com/bitebait/curry/config"
	"github.com/bitebait/curry/crawler"
	"github.com/bitebait/curry/scheduler"
)

func main() {
	wg := sync.WaitGroup{}

	initApi(&wg)
	initCrawler(&wg)

	wg.Wait()
}

func initApi(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()

		cfg := config.GetConfig()

		db.Init(cfg.DB.Name)

		routes.Init()

		fmt.Println(cfg.App.AsciiName)
		log.Printf("Running and Listening on %s:%s\n", cfg.App.Host, cfg.App.Port)
		log.Printf("API Endpoint: %s\n", cfg.Api.Endpoint)
		log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.App.Host, cfg.App.Port), nil))
	}()

}

func initCrawler(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()

		s := scheduler.Init(
			func() {
				log.Println("Running crawler...")
				cache := &models.Cache{}
				dB := db.GetDB()
				cache.Stores = *crawler.Run()
				dB.Create(cache)
			},
		)

		<-s.Start()
	}()

}
