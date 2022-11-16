package spiders

import (
	"github.com/gocolly/colly"
)

func init() {
	NewSpider(
		"atacadocollections",
		".price",
		"https://www.atacadocollections.com/",
		func(e *colly.HTMLElement) string {
			return e.Text
		},
	)
}
