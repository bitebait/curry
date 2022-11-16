package spiders

import (
	"strings"

	"github.com/gocolly/colly"
)

func init() {
	NewSpider(
		"casaamericana",
		"span.justify-center:nth-child(2)",
		"https://www.casaamericana.com.py/",
		func(e *colly.HTMLElement) string {
			return strings.Split(e.Text, "R$")[1]
		},
	)
}
