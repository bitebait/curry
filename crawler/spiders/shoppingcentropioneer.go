package spiders

import (
	"strings"

	"github.com/bitebait/curry/api/models"
	"github.com/bitebait/curry/config"

	"github.com/gocolly/colly"
)

func ShoppingCentro(channel chan models.Store) {
	spider := &config.Spider{}
	spider.Name = "shoppingcentropioneer"
	spider.Channel = channel
	spider.Selector = "div.h-4:nth-child(1)"
	spider.URL = "https://shoppingcentropioneer.com/"
	spider.GetValue = func(e *colly.HTMLElement) string {
		data := strings.TrimSpace(strings.ReplaceAll(e.Text, ".", ""))
		return data
	}

	runSpider(spider)
}
