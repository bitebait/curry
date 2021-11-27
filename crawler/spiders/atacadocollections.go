package spiders

import (
	"curry/api/models"
	"curry/config"
	"strings"

	"github.com/gocolly/colly"
)

func AtacadoCollections(channel chan models.Store) {
	spider := &config.Spider{}
	spider.Name = "atacadocollections"
	spider.Channel = channel
	spider.Selector = ".price"
	spider.URL = "https://www.atacadocollections.com/"
	spider.GetValue = func(e *colly.HTMLElement) string {
		data := strings.TrimSpace(e.Text)
		return data
	}

	runSpider(spider)
}
