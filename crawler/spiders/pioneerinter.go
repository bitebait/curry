package spiders

import (
	"curry/api/models"
	"curry/config"
	"strings"

	"github.com/gocolly/colly"
)

func PioneerInter(channel chan models.Store) {
	spider := &config.Spider{}
	spider.Name = "pioneerinter"
	spider.Channel = channel
	spider.Selector = "li.dropDown:nth-child(2) > nav:nth-child(2) > a:nth-child(1)"
	spider.URL = "https://www.pioneerinter.com/"
	spider.GetValue = func(e *colly.HTMLElement) string {
		data := strings.TrimSpace(strings.Split(e.Text, "=")[1])
		return data
	}

	runSpider(spider)
}
