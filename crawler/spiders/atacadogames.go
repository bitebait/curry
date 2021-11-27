package spiders

import (
	"github.com/bitebait/curry/api/models"
	"github.com/bitebait/curry/config"
	"strings"

	"github.com/gocolly/colly"
)

func AtacadoGames(channel chan models.Store) {
	spider := &config.Spider{}
	spider.Name = "atacadogames"
	spider.Channel = channel
	spider.Selector = "div.header-extra div.wrap.is-centralized ul.grid li.grid-item span.text.has-icon"
	spider.URL = "https://www.atacadogames.com/"
	spider.GetValue = func(e *colly.HTMLElement) string {
		data := strings.TrimSpace(strings.Split(e.Text, "Cotação")[1])
		return data
	}

	runSpider(spider)
}
