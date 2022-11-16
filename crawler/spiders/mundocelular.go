package spiders

import (
	"github.com/gocolly/colly"
)

func init() {
	NewSpider(
		"mundodocelular",
		".heading > div:nth-child(1) > span:nth-child(4)",
		"https://www.mundodocelular.com/",
		func(e *colly.HTMLElement) string {
			return e.Text
		},
	)
}
