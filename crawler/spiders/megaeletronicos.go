package spiders

import (
	"github.com/gocolly/colly"
)

func init() {
	NewSpider(
		"megaeletronicos",
		"div.d-flex.flex-row.align-items-center.cotizacion p:nth-child(4)",
		"https://www.megaeletronicos.com/br",
		func(e *colly.HTMLElement) string {
			return e.Text
		},
	)
}
