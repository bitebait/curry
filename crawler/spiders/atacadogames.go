package spiders

import (
	"strings"

	"github.com/gocolly/colly"
)

func init() {
	NewSpider(
		"atacadogames",
		"div.header-extra div.wrap.is-centralized ul.grid li.grid-item span.text.has-icon",
		"https://www.atacadogames.com/",
		func(e *colly.HTMLElement) string {
			return strings.Split(e.Text, "Cotação")[1]
		},
	)
}
