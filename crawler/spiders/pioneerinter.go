package spiders

import (
	"strings"

	"github.com/gocolly/colly"
)

func init() {
	NewSpider(
		"pioneerinter",
		".currency-selector > span:nth-child(1)",
		"https://www.pioneerinter.com/",
		func(e *colly.HTMLElement) string {
			return strings.Split(e.Text, "=")[1]
		},
	)
}
