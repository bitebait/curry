package main

import (
	"fmt"
	"sync"

	"github.com/bitebait/curry/api"
	"github.com/bitebait/curry/crawler"
)

func main() {
	fmt.Print("teste")
	wg := sync.WaitGroup{}
	fmt.Print("teste2")
	api.Init(&wg)
	fmt.Print("teste3")
	crawler.Init(&wg)
	fmt.Print("teste4")

	wg.Wait()
	fmt.Print("teste5")
}
