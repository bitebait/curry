package spiders

import (
	"github.com/gocolly/colly"
)

var EleganciaCompany = &Spider{
	Name:     "eleganciacompany",
	Selector: "div.quotation span",
	URL:      "https://www.eleganciacompany.com/",
	GetValue: func(e *colly.HTMLElement) string {
		return e.Text
	},
}
