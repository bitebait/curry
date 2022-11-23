package spiders

import (
	"strings"

	"github.com/gocolly/colly"
)

func init() {
	NewSpider(
		"aromastore",
		".topbar > div:nth-child(1) > div:nth-child(1) > div:nth-child(1) > div:nth-child(1)",
		"https://www.aromastore.com.br/",
		func(e *colly.HTMLElement) string {
			return strings.Split(e.Text, "R$")[1]
		},
	)
}
