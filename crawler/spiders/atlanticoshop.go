package spiders

import (
	"github.com/gocolly/colly"
)

func init() {
	NewSpider(
		"atlanticoshop",
		"div.bandera:nth-child(3) > small:nth-child(2)",
		"https://www.atlanticoshop.com.br/",
		func(e *colly.HTMLElement) string {
			return e.Text
		},
	)
}
