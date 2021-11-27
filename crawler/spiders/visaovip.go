package spiders

import (
	"curry/api/models"
	"curry/config"
	"strings"

	"github.com/gocolly/colly"
)

func VisaoVip(channel chan models.Store) {
	spider := &config.Spider{}
	spider.Name = "visaovip"
	spider.Channel = channel
	spider.Selector = "div.ui-g-12 div.ui-g-6.painel-cambio-topo:nth-child(1)"
	spider.URL = "http://www.visaovip.com/"
	spider.GetValue = func(e *colly.HTMLElement) string {
		data := strings.TrimSpace(e.Text)
		return data
	}

	runSpider(spider)
}
