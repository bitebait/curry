package spiders

import (
	"github.com/gocolly/colly"
)

func init() {
	NewSpider(
		"dolarpy",
		".blue2",
		"https://www.dolarpy.com.br/",
		func(e *colly.HTMLElement) string {
			return e.Text
		},
	)
}
