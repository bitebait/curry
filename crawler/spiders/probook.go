package spiders

import (
	"github.com/gocolly/colly"
	"strings"
)

func init() {
	NewSpider(
		"probook",
		".text-header-foreground",
		"https://www.probook.com.py/",
		func(e *colly.HTMLElement) string {
			return strings.Split(e.Text, "x")[1]
		},
	)
}
