package spiders

import (
	"github.com/gocolly/colly"
	"strings"
)

func init() {
	NewSpider(
		"madridcenter",
		"div.col-lg-4:nth-child(1) > div:nth-child(1) > div:nth-child(1)",
		"https://www.madridcenter.com/",
		func(e *colly.HTMLElement) string {
			return strings.Split(e.Text, " ")[3]
		},
	)
}
