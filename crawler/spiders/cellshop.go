package spiders

import (
	"strings"

	"github.com/gocolly/colly"
)

func init() {
	NewSpider(
		"cellshop",
		".showCur",
		"https://www.cellshop.com/br/",
		func(e *colly.HTMLElement) string {
			return strings.Split(e.Text, "Câmbio")[1]
		},
	)
}
