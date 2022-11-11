package spiders

import (
	"strings"

	"github.com/gocolly/colly"
)

var EleganciaCompany = &Spider{
	Name:     "eleganciacompany",
	Selector: "div.quotation span",
	URL:      "https://www.eleganciacompany.com/",
	GetValue: func(e *colly.HTMLElement) string {
		data := strings.TrimSpace(e.Text)
		return data
	},
}
