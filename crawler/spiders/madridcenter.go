package spiders

import (
	"strings"

	"github.com/bitebait/curry/api/models"
	"github.com/bitebait/curry/config"

	"github.com/gocolly/colly"
)

func MadridCenter(channel chan models.Store) {
	spider := &config.Spider{}
	spider.Name = "madridcenter"
	spider.Channel = channel
	spider.Selector = "body > header > div > div > div > div.item.top-quotes > div.col-12.text3 > span > span:nth-child(1) > strong"
	spider.URL = "https://www.madridcenter.com/"
	spider.GetValue = func(e *colly.HTMLElement) string {
		data := strings.TrimSpace(e.Text)
		return data

	}

	runSpider(spider)
}
