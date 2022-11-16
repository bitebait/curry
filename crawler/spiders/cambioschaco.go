package spiders

import (
	"github.com/gocolly/colly"
)

func init() {
	NewSpider(
		"cambioschaco",
		"#arb-exchange-brl > td:nth-child(3) > span:nth-child(1)",
		"https://www.cambioschaco.com.py/pt-br/",
		func(e *colly.HTMLElement) string {
			return e.Text
		},
	)
}
