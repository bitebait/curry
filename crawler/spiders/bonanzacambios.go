package spiders

import (
	"strings"

	"github.com/bitebait/curry/api/models"
	"github.com/bitebait/curry/config"

	"github.com/gocolly/colly"
)

func BonanzaCambios(channel chan models.Store) {
	spider := &config.Spider{}
	spider.Name = "bonanzacambios"
	spider.Channel = channel
	spider.Selector = "div.content-inner:nth-child(1) > div:nth-child(2) > div:nth-child(2) > section:nth-child(1) > div:nth-child(1) > div:nth-child(1) > div:nth-child(1) > div:nth-child(1) > table:nth-child(1) > tbody:nth-child(2) > tr:nth-child(3) > td:nth-child(5)"
	spider.URL = "https://bonanzacambios.com.py/"
	spider.GetValue = func(e *colly.HTMLElement) string {
		data := strings.TrimSpace(e.Text)
		return data
	}

	runSpider(spider)
}
