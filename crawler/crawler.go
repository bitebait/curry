package crawler

import (
	"log"
	"sync"

	"github.com/bitebait/curry/api/db"
	"github.com/bitebait/curry/api/models"
	"github.com/bitebait/curry/crawler/spiders"
	"github.com/bitebait/curry/scheduler"
)

func Init(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()

		s := scheduler.Init(
			func() {
				log.Println("Running crawler...")
				cache := &models.Cache{}
				dB := db.GetDB()
				cache.Stores = *runCrawler()
				dB.Create(cache)
			},
		)

		<-s.Start()
	}()

}

func runCrawler() *[]models.Store {
	var stores []models.Store

	wg := sync.WaitGroup{}
	channel := make(chan models.Store)

	for _, spider := range spiders.AllSpiders {
		wg.Add(1)
		go func(spider spiders.Runnable) {
			defer wg.Done()
			spider.RunSpider(channel)
		}(spider)
	}

	go func() {
		wg.Wait()
		close(channel)
	}()

	for i := range channel {
		stores = append(stores, i)
	}

	log.Printf("Success: %v urls visited.", len(stores))

	return &stores
}
