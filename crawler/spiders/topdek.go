package spiders

import (
	"strings"

	"github.com/gocolly/colly"
)

func init() {
	NewSpider(
		"topdek",
		"body > div:nth-child(8) > div:nth-child(1) > p:nth-child(1)",
		"https://www.topdek.com.br/br",
		func(e *colly.HTMLElement) string {
			return strings.Split(e.Text, " ")[3]
		},
	)
}
