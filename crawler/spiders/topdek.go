package spiders

import (
	"strings"

	"github.com/bitebait/curry/api/models"
	"github.com/bitebait/curry/config"

	"github.com/gocolly/colly"
)

func TopDek(channel chan models.Store) {
	spider := &config.Spider{}
	spider.Name = "topdek"
	spider.Channel = channel
	spider.Selector = "body > div:nth-child(8) > div:nth-child(1) > p:nth-child(1)"
	spider.URL = "https://www.topdek.com.br/br"
	spider.GetValue = func(e *colly.HTMLElement) string {
		data := strings.TrimSpace(strings.Split(e.Text, " ")[3])
		return data
	}

	runSpider(spider)
}
