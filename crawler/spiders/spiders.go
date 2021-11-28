package spiders

import (
	"log"
	"net/http"

	"github.com/bitebait/curry/api/models"
	"github.com/bitebait/curry/config"
	"github.com/bitebait/curry/utils"

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
		AlboradaInfo,
		OneClick,
		TeCombras,
	}
}

func runSpider(spider *config.Spider) {
	c := colly.NewCollector(
		colly.Async(true),
	)

	c.WithTransport(&http.Transport{
		DisableKeepAlives: true,
	})

	cfg := config.GetConfig()

	c.OnHTML(spider.Selector, func(e *colly.HTMLElement) {
		store := &models.Store{
			Name:     spider.Name,
			Currency: cfg.Currency.Currency,
			Value:    utils.FormatCurrency(spider.GetValue(e)),
			URL:      spider.URL,
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
