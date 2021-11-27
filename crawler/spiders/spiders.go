package spiders

import (
	"github.com/bitebait/curry/api/models"
	"github.com/bitebait/curry/config"
	"github.com/bitebait/curry/utils"
	"log"
	"net/http"

	"github.com/gocolly/colly"
)

func AllSpiders() []func(channel chan models.Store) {
	return []func(channel chan models.Store){
		CellShop,
		VictoriaStore,
		PioneerInter,
		MadridCenter,
		MegaEletronicos,
		AtacadoGames,
		MobileZone,
		VisaoVip,
		EleganciaCompany,
		ShoppingCentro,
		TopDek,
		AtacadoCollections,
		AtlanticoShop,
		AudiumEletronics,
	}
}

func runSpider(spider *config.Spider) {
	c := colly.NewCollector(
		colly.Async(true),
	)

	c.WithTransport(&http.Transport{
		DisableKeepAlives: true,
	})

	c.OnHTML(spider.Selector, func(e *colly.HTMLElement) {
		store := &models.Store{
			Name:  spider.Name,
			Value: utils.FormatCurrency(spider.GetValue(e)),
			URL:   spider.URL,
		}
		spider.Channel <- *store
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	err := c.Visit(spider.URL)
	if err != nil {
		log.Panicf("Failed to visit URL: %s", spider.URL)
	}
	c.Wait()
}
