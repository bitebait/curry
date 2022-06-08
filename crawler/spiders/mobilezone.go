package spiders

import (
	"strings"

	"github.com/bitebait/curry/api/models"
	"github.com/bitebait/curry/config"

	"github.com/gocolly/colly"
)

func MobileZone(channel chan models.Store) {
	spider := &config.Spider{}
	spider.Name = "mobilezone"
	spider.Channel = channel
	spider.Selector = "div.bandera:nth-child(2) > small:nth-child(2)"
	spider.URL = "https://www.mobilezone.com.br/"
	spider.GetValue = func(e *colly.HTMLElement) string {
		data := strings.TrimSpace(e.Text)
		return data
	}

	runSpider(spider)
}
