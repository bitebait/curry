package spiders

import (
	"strings"

	"github.com/bitebait/curry/api/models"
	"github.com/bitebait/curry/config"

	"github.com/gocolly/colly"
)

func VisaoVip(channel chan models.Store) {
	spider := &config.Spider{}
	spider.Name = "visaovip"
	spider.Channel = channel
	spider.Selector = "div.ui-panelgrid-cell:nth-child(1) > span:nth-child(1)"
	spider.URL = "http://visaovip.com/"
	spider.GetValue = func(e *colly.HTMLElement) string {
		data := strings.TrimSpace(e.Text)
		return data
	}

	runSpider(spider)
}
