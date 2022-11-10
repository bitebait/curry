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
	"github.com/bitebait/curry/utils"
)

func AllSpiders() []func(channel chan models.Store) {
	return []func(channel chan models.Store){
		AlboradaInfo,
		AtacadoCollections,
		AtacadoGames,
		AtlanticoShop,
		AudiumEletronics,
		BonanzaCambios,
		CambiosChaco,
		CellShop,
		DolarPy,
		EleganciaCompany,
		MadridCenter,
		MegaEletronicos,
		MercosurCambios,
		MobileZone,
		OneClick,
		PioneerInter,
		ShoppingCentro,
		TeCombras,
		TopDek,
		VictoriaStore,
		VisaoVip,
	}
}

func runSpider(spider *config.Spider) {
	cfg := config.GetConfig()

	c := colly.NewCollector(
		colly.Async(true),
	)

	extensions.RandomUserAgent(c)
	c.WithTransport(&http.Transport{
		DisableKeepAlives:     true,
		ResponseHeaderTimeout: time.Duration(cfg.Crawler.ResponseHeaderTimeout) * time.Second,
	})

	c.SetRequestTimeout(time.Duration(cfg.Crawler.ClientTimeout) * time.Second)

	c.OnHTML(spider.Selector, func(e *colly.HTMLElement) {
		store := &models.Store{
			Name:     spider.Name,
			Currency: cfg.Currency.Currency,
			Value:    utils.FormatCurrency(spider.GetValue(e)),
			URL:      spider.URL,
		}
		spider.Channel <- *store
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

	err := c.Visit(spider.URL)
	if err != nil {
		log.Panicf("FATAL: Failed to visit URL: %s\n", spider.URL)
	}

	c.Wait()
}
