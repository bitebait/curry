package crawler

import (
	"log"
	"sync"

	"github.com/bitebait/curry/api/models"
	"github.com/bitebait/curry/crawler/spiders"
)

func Run() *[]models.Store {
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
