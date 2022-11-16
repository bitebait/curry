package spiders

import (
	"github.com/gocolly/colly"
)

func init() {
	NewSpider(
		"tcheloco",
		".nohover",
		"https://www.tcheloco.com.py/br/",
		func(e *colly.HTMLElement) string {
			return e.Text
		},
	)
}
