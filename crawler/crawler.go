package crawler

import (
	"log"
	"sync"
	"time"

	"github.com/bitebait/curry/api/models"
	"github.com/bitebait/curry/cache"
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
				cache.SetCache(&models.Cache{
					Stores: *runCrawler(),
				})

			},
		)
		<-s.Start()
	}()
}

func runCrawler() *[]models.Store {
	var stores []models.Store
	var allSpiders = *spiders.GetAllSpiders()

	now := time.Now()
	defer func() {
		log.Printf("CRAWLER function took %s.", time.Since(now))
	}()

	wg := sync.WaitGroup{}
	channel := make(chan models.Store)

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

	for i := range channel {
		stores = append(stores, i)
	}

	log.Printf("FINISHED: %v of %v urls visited.", len(stores), len(allSpiders))

	return &stores
}
