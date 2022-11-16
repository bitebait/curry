package spiders

import (
	"strings"

	"github.com/gocolly/colly"
)

func init() {
	NewSpider(
		"tecombras",
		".currency-selector > span:nth-child(1)",
		"https://www.tecombras.net/",
		func(e *colly.HTMLElement) string {
			return strings.Split(e.Text, "=")[1]
		},
	)
}
