package spiders

import (
	"github.com/gocolly/colly"
)

func init() {
	NewSpider(
		"lgimportados",
		".quotation > span:nth-child(1) > strong:nth-child(1)",
		"https://www.lgimportados.com/",
		func(e *colly.HTMLElement) string {
			return e.Text
		},
	)
}
