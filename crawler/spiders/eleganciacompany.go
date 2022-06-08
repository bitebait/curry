package spiders

import (
	"strings"

	"github.com/bitebait/curry/api/models"
	"github.com/bitebait/curry/config"

	"github.com/gocolly/colly"
)

func EleganciaCompany(channel chan models.Store) {
	spider := &config.Spider{}
	spider.Name = "eleganciacompany"
	spider.Channel = channel
	spider.Selector = "div.quotation span"
	spider.URL = "https://www.eleganciacompany.com/"
	spider.GetValue = func(e *colly.HTMLElement) string {
		data := strings.TrimSpace(e.Text)
		return data
	}

	runSpider(spider)
}
