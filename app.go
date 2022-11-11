package main

import (
	"sync"

	"github.com/bitebait/curry/api"
	"github.com/bitebait/curry/crawler"
)

func main() {
	wg := sync.WaitGroup{}

	api.Init(&wg)
	crawler.Init(&wg)

	wg.Wait()
}
