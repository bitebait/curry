package spiders

import (
	"strings"

	"github.com/bitebait/curry/api/models"
	"github.com/bitebait/curry/config"

	"github.com/gocolly/colly"
)

func PontoCom(channel chan models.Store) {
	spider := &config.Spider{}
	spider.Name = "pontocom"
	spider.Channel = channel
	spider.Selector = ".color_red"
	spider.URL = "https://www.pontocom.com/"
	spider.GetValue = func(e *colly.HTMLElement) string {
		data := strings.TrimSpace(e.Text)
		return data
	}

	runSpider(spider)
}
