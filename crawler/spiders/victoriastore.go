package spiders

import (
	"strings"

	"github.com/gocolly/colly"
)

func init() {
	NewSpider(
		"victoriastore",
		".left-content > div:nth-child(1) > ul:nth-child(1) > li:nth-child(2) > div:nth-child(1) > span:nth-child(1)",
		"https://www.victoriastore.com.br/",
		func(e *colly.HTMLElement) string {
			return strings.Split(e.Text, "=")[1]
		},
	)
}
