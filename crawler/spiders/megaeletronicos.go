package spiders

import (
	"strings"

	"github.com/gocolly/colly"
)

var MegaEletronicos = &Spider{
	Name:     "megaeletronicos",
	Selector: "div.d-flex.flex-row.align-items-center.cotizacion p:nth-child(4)",
	URL:      "https://www.megaeletronicos.com/br",
	GetValue: func(e *colly.HTMLElement) string {
		data := strings.TrimSpace(e.Text)
		return data
	},
}
