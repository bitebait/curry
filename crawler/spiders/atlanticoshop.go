package spiders

import (
	"github.com/gocolly/colly"
)

var atlanticoShop = &Spider{
	Name:     "atlanticoshop",
	Selector: "div.bandera:nth-child(3) > small:nth-child(2)",
	URL:      "https://www.atlanticoshop.com.br/",
	GetValue: func(e *colly.HTMLElement) string {
		return e.Text
	},
}
