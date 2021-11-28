package spiders

import (
	"strings"

	"github.com/bitebait/curry/api/models"
	"github.com/bitebait/curry/config"

	"github.com/gocolly/colly"
)

func AlboradaInfo(channel chan models.Store) {
	spider := &config.Spider{}
	spider.Name = "alboradainfo"
	spider.Channel = channel
	spider.Selector = ".quotation > span:nth-child(2)"
	spider.URL = "https://www.alboradainfo.com/"
	spider.GetValue = func(e *colly.HTMLElement) string {
		data := strings.TrimSpace(e.Text)
		return data
	}

	runSpider(spider)
}
