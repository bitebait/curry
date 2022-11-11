package spiders

import (
	"strings"

	"github.com/gocolly/colly"
)

var AtlanticoShop = &Spider{
	Name:     "atlanticoshop",
	Selector: "div.bandera:nth-child(3) > small:nth-child(2)",
	URL:      "https://www.atlanticoshop.com.br/",
	GetValue: func(e *colly.HTMLElement) string {
		data := strings.TrimSpace(e.Text)
		return data
	},
}
