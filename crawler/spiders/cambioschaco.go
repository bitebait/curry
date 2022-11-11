package spiders

import (
	"strings"

	"github.com/gocolly/colly"
)

var CambiosChaco = &Spider{
	Name:     "cambioschaco",
	Selector: "#arb-exchange-brl > td:nth-child(3) > span:nth-child(1)",
	URL:      "https://www.cambioschaco.com.py/pt-br/",
	GetValue: func(e *colly.HTMLElement) string {
		data := strings.TrimSpace(e.Text)
		return data
	},
}
