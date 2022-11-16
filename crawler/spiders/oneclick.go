package spiders

import (
	"github.com/gocolly/colly"
)

func init() {
	NewSpider(
		"oneclick",
		".login-box > center:nth-child(1) > h5:nth-child(1)",
		"https://oneclick.com.py/",
		func(e *colly.HTMLElement) string {
			return e.Text
		},
	)
}
