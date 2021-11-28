package spiders

import (
	"strings"

	"github.com/bitebait/curry/api/models"
	"github.com/bitebait/curry/config"

	"github.com/gocolly/colly"
)

func OneClick(channel chan models.Store) {
	spider := &config.Spider{}
	spider.Name = "oneclick"
	spider.Channel = channel
	spider.Selector = ".login-box > center:nth-child(1) > h5:nth-child(1)"
	spider.URL = "https://oneclick.com.py/"
	spider.GetValue = func(e *colly.HTMLElement) string {
		data := strings.TrimSpace(e.Text)
		return data
	}

	runSpider(spider)
}
