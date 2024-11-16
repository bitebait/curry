package spiders

import (
	"github.com/gocolly/colly"
	"strings"
)

func init() {
	NewSpider(
		"lgimportados",
		"div.is-cambio-us > div.ms-n3",
		"https://www.lgimportados.com/",
		func(e *colly.HTMLElement) string {
			return strings.Split(e.Text, ":")[1]
		},
	)
}
