package spiders

import (
	"github.com/gocolly/colly"
)

var cambiosChaco = &Spider{
	Name:     "cambioschaco",
	Selector: "#arb-exchange-brl > td:nth-child(3) > span:nth-child(1)",
	URL:      "https://www.cambioschaco.com.py/pt-br/",
	GetValue: func(e *colly.HTMLElement) string {
		return e.Text
	},
}
