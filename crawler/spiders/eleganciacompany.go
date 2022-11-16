package spiders

import (
	"github.com/gocolly/colly"
)

func init() {
	NewSpider(
		"eleganciacompany",
		"div.quotation span",
		"https://www.eleganciacompany.com/",
		func(e *colly.HTMLElement) string {
			return e.Text
		},
	)
}
