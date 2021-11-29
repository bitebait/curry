package spiders

import (
	"strings"

	"github.com/bitebait/curry/api/models"
	"github.com/bitebait/curry/config"

	"github.com/gocolly/colly"
)

func CambiosChaco(channel chan models.Store) {
	spider := &config.Spider{}
	spider.Name = "cambioschaco"
	spider.Channel = channel
	spider.Selector = "#arb-exchange-brl > td:nth-child(3) > span:nth-child(1)"
	spider.URL = "https://www.cambioschaco.com.py/pt-br/"
	spider.GetValue = func(e *colly.HTMLElement) string {
		data := strings.TrimSpace(e.Text)
		return data
	}

	runSpider(spider)
}
