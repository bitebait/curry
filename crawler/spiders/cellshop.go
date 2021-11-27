package spiders

import (
	"curry/api/models"
	"curry/config"
	"strings"

	"github.com/gocolly/colly"
)

func CellShop(channel chan models.Store) {
	spider := config.Spider{}
	spider.Name = "cellshop"
	spider.Channel = channel
	spider.Selector = ".showCur"
	spider.URL = "https://www.cellshop.com/br/"
	spider.GetValue = func(e *colly.HTMLElement) string {
		data := strings.TrimSpace(strings.Split(e.Text, "Câmbio")[1])
		return data
	}

	runSpider(&spider)
}
