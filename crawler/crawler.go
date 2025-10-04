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
		schedulerInstance := scheduler.Init(runCrawler)
		<-schedulerInstance.Start()
	}()
}

func runCrawler() {
	allSpiders := spiders.GetAllSpiders()
	storeChannel := make(chan models.Store, len(allSpiders))
	var stores []models.Store
	var wg sync.WaitGroup

	for _, spider := range allSpiders {
		wg.Add(1)
		go runSpider(spider, storeChannel, &wg)
	}

	go closeChannelWhenDone(&wg, storeChannel)

	for store := range storeChannel {
		stores = append(stores, store)
	}

	cache.SetCache(&models.Cache{Stores: stores})
	log.Printf("FINISHED: %v of %v URLs visited.", len(stores), len(allSpiders))
}

func runSpider(spider spiders.Runnable, storeChannel chan models.Store, wg *sync.WaitGroup) {
	defer wg.Done()
	spider.RunSpider(storeChannel)
}

func closeChannelWhenDone(wg *sync.WaitGroup, storeChannel chan models.Store) {
	wg.Wait()
	close(storeChannel)
}
