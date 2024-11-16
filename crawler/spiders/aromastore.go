package spiders

import (
	"strings"

	"github.com/gocolly/colly"
)

func init() {
	NewSpider(
		"aromastore",
		"div.quotation",
		"https://www.aromastore.com.py",
		func(e *colly.HTMLElement) string {
			return strings.Split(e.Text, "x")[1]
		},
	)
}
