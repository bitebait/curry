package spiders

import (
	"strings"

	"github.com/bitebait/curry/api/models"
	"github.com/bitebait/curry/config"

	"github.com/gocolly/colly"
)

func AudiumEletronics(channel chan models.Store) {
	spider := &config.Spider{}
	spider.Name = "audiumelectronics"
	spider.Channel = channel
	spider.Selector = "div.quotation:nth-child(3) > span:nth-child(1)"
	spider.URL = "https://www.audiumelectronics.com/home"
	spider.GetValue = func(e *colly.HTMLElement) string {
		data := strings.TrimSpace(e.Text)
		return data
	}

	runSpider(spider)
}
