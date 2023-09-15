package crawler

import (
	"log"
	"sync"

	"github.com/bitebait/curry/api/models"
	"github.com/bitebait/curry/cache"
	"github.com/bitebait/curry/crawler/spiders"
	"github.com/bitebait/curry/scheduler"
)

func Init(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()

		s := scheduler.Init(runCrawler)
		<-s.Start()
	}()
}

func runCrawler() {
	allSpiders := spiders.GetAllSpiders()
	var stores []models.Store
	var wg sync.WaitGroup
	channel := make(chan models.Store, len(allSpiders))

	for _, spider := range allSpiders {
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

	for store := range channel {
		stores = append(stores, store)
	}

	cache.SetCache(&models.Cache{
		Stores: stores,
	})

	log.Printf("FINISHED: %v of %v urls visited.", len(stores), len(allSpiders))
}
