package spiders

import (
	"curry/api/models"
	"curry/config"
	"strings"

	"github.com/gocolly/colly"
)

func VictoriaStore(channel chan models.Store) {
	spider := &config.Spider{}
	spider.Name = "victoriastore"
	spider.Channel = channel
	spider.Selector = "div.currency-selector"
	spider.URL = "https://www.victoriastore.com.br/"
	spider.GetValue = func(e *colly.HTMLElement) string {
		data := strings.TrimSpace(strings.Split(e.ChildText("span"), "=")[1])
		return data
	}

	runSpider(spider)
}
