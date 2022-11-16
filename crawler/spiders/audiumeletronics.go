package spiders

import (
	"github.com/gocolly/colly"
)

func init() {
	NewSpider(
		"audiumelectronics",
		"div.quotation:nth-child(3) > span:nth-child(1)",
		"https://www.audiumelectronics.com/home",
		func(e *colly.HTMLElement) string {
			return e.Text
		},
	)
}
