package crawler

import (
	"github.com/bitebait/curry/api/models"
	"github.com/bitebait/curry/crawler/spiders"
	"log"
	"sync"
)

func Run() *[]models.Store {
	var stores []models.Store

	wg := sync.WaitGroup{}
	channel := make(chan models.Store)

	for _, spider := range spiders.AllSpiders() {
		runParallel(spider, &wg, channel)
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

func runParallel(f func(c chan models.Store), wg *sync.WaitGroup, c chan models.Store) {
	wg.Add(1)

	go func() {
		defer wg.Done()
		f(c)
	}()
}
