package spiders

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"

	"github.com/bitebait/curry/api/models"
	"github.com/bitebait/curry/config"
	"github.com/bitebait/curry/helpers"
)

type Spider struct {
	Name     string
	Selector string
	GetValue func(e *colly.HTMLElement) string
	URL      string
}

type Runnable interface {
	RunSpider(channel chan models.Store)
}

func (s Spider) RunSpider(channel chan models.Store) {

	c := colly.NewCollector(
		colly.Async(true),
	)

	extensions.RandomUserAgent(c)
	c.WithTransport(&http.Transport{
		DisableKeepAlives:     true,
		ResponseHeaderTimeout: time.Duration(config.GetConfig.Crawler.ResponseHeaderTimeout) * time.Second,
	})

	c.SetRequestTimeout(time.Duration(config.GetConfig.Crawler.ClientTimeout) * time.Second)

	c.OnHTML(s.Selector, func(e *colly.HTMLElement) {
		store := &models.Store{
			Name:     s.Name,
			Currency: config.GetConfig.Currency.Currency,
			Value:    helpers.FormatCurrency(s.GetValue(e)),
			URL:      s.URL,
		}

		channel <- *store
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("SUCCESS: Visited", r.Request.URL)
	})
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("SUCCESS: Scrap Finished", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Printf("FATAL: Something went wrong (%s): %s", r.Request.URL, err)
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("INFO: Visiting", r.URL)
	})

	err := c.Visit(s.URL)
	if err != nil {
		log.Panicf("FATAL: Failed to visit URL: %s\n", s.URL)
	}

	c.Wait()
}

var AllSpiders []Runnable = []Runnable{
	alboradaInfo,
	atacadoCollections,
	atacadoGames,
	atlanticoShop,
	audiumEletronics,
	bonanzaCambios,
	cambiosChaco,
	casaAmericana,
	cellShop,
	comprasParaguai,
	dolarPy,
	eleganciaCompany,
	gabbaHobby,
	hbGames,
	icompy,
	lgImportados,
	madridCenter,
	megaEletro,
	megaEletronicos,
	mercosurCambios,
	mundoCelular,
	oneClick,
	pioneerInter,
	proBook,
	shoppingCentro,
	tcheLoco,
	teCombras,
	topDek,
	victoriaStore,
	visaoVip,
}
