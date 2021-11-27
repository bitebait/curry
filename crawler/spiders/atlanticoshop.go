package spiders

import (
	"github.com/bitebait/curry/api/models"
	"github.com/bitebait/curry/config"
	"strings"

	"github.com/gocolly/colly"
)

func AtlanticoShop(channel chan models.Store) {
	spider := &config.Spider{}
	spider.Name = "atlanticoshop"
	spider.Channel = channel
	spider.Selector = "div.bandera:nth-child(3) > small:nth-child(2)"
	spider.URL = "https://www.atlanticoshop.com.br/"
	spider.GetValue = func(e *colly.HTMLElement) string {
		data := strings.TrimSpace(e.Text)
		return data
	}

	runSpider(spider)
}
