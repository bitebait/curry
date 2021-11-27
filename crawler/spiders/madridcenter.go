package spiders

import (
	"github.com/bitebait/curry/api/models"
	"github.com/bitebait/curry/config"
	"strings"

	"github.com/gocolly/colly"
)

func MadridCenter(channel chan models.Store) {
	spider := &config.Spider{}
	spider.Name = "madridcenter"
	spider.Channel = channel
	spider.Selector = "body > header:nth-child(4) > div:nth-child(2) > div:nth-child(1) > div:nth-child(1) > div:nth-child(5) > div:nth-child(3) > span:nth-child(1) > span:nth-child(1) > strong:nth-child(2)"
	spider.URL = "https://www.madridcenter.com/"
	spider.GetValue = func(e *colly.HTMLElement) string {
		data := strings.TrimSpace(e.Text)
		return data

	}

	runSpider(spider)
}
