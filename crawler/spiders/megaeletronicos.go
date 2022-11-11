package spiders

import (
	"github.com/gocolly/colly"
)

var MegaEletronicos = &Spider{
	Name:     "megaeletronicos",
	Selector: "div.d-flex.flex-row.align-items-center.cotizacion p:nth-child(4)",
	URL:      "https://www.megaeletronicos.com/br",
	GetValue: func(e *colly.HTMLElement) string {
		return e.Text
	},
}
