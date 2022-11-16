package spiders

import (
	"github.com/gocolly/colly"
)

func init() {
	NewSpider(
		"madridcenter",
		"body > header > div > div > div > div.item.top-quotes > div.col-12.text3 > span > span:nth-child(1) > strong",
		"https://www.madridcenter.com/",
		func(e *colly.HTMLElement) string {
			return e.Text
		},
	)
}
