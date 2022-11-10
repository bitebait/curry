package spiders

import (
	"strings"

	"github.com/bitebait/curry/api/models"
	"github.com/bitebait/curry/config"

	"github.com/gocolly/colly"
)

func VictoriaStore(channel chan models.Store) {
	spider := &config.Spider{}
	spider.Name = "victoriastore"
	spider.Channel = channel
	spider.Selector = ".left-content > div:nth-child(1) > ul:nth-child(1) > li:nth-child(2) > div:nth-child(1) > span:nth-child(1)"
	spider.URL = "https://www.victoriastore.com.br/"
	spider.GetValue = func(e *colly.HTMLElement) string {
		data := strings.TrimSpace(strings.Split(e.Text, "=")[1])
		return data
	}

	runSpider(spider)
}
