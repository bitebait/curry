package spiders

import (
	"github.com/gocolly/colly"
)

func init() {
	NewSpider(
		"probook",
		".quotation > span:nth-child(1) > strong:nth-child(1)",
		"https://www.probook.com.py/",
		func(e *colly.HTMLElement) string {
			return e.Text
		},
	)
}
