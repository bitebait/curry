package spiders

import (
	"strings"

	"github.com/bitebait/curry/api/models"
	"github.com/bitebait/curry/config"

	"github.com/gocolly/colly"
)

func TeCombras(channel chan models.Store) {
	spider := &config.Spider{}
	spider.Name = "tecombras"
	spider.Channel = channel
	spider.Selector = ".currency-selector > span:nth-child(1)"
	spider.URL = "https://www.tecombras.net/"
	spider.GetValue = func(e *colly.HTMLElement) string {
		data := strings.TrimSpace(strings.Split(e.Text, "=")[1])
		return data
	}

	runSpider(spider)
}
