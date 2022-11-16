package spiders

import (
	"github.com/gocolly/colly"
)

func init() {
	NewSpider(
		"megaeletro",
		"div.mr-7:nth-child(3) > p:nth-child(2)",
		"https://www.megaeletro.com.py/br",
		func(e *colly.HTMLElement) string {
			return e.Text
		},
	)
}
