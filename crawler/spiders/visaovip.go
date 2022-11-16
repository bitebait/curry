package spiders

import (
	"github.com/gocolly/colly"
)

func init() {
	NewSpider(
		"visaovip",
		"div.ui-panelgrid-cell:nth-child(1) > span:nth-child(1)",
		"http://visaovip.com/",
		func(e *colly.HTMLElement) string {
			return e.Text
		},
	)
}
