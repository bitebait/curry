package spiders

import (
	"github.com/bitebait/curry/api/models"
	"github.com/bitebait/curry/config"
	"strings"

	"github.com/gocolly/colly"
)

func MegaEletronicos(channel chan models.Store) {
	spider := &config.Spider{}
	spider.Name = "megaeletronicos"
	spider.Channel = channel
	spider.Selector = "div.d-flex.flex-row.align-items-center.cotizacion p:nth-child(4)"
	spider.URL = "https://www.megaeletronicos.com/br"
	spider.GetValue = func(e *colly.HTMLElement) string {
		data := strings.TrimSpace(e.Text)
		return data
	}

	runSpider(spider)
}
