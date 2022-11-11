package spiders

import (
	"github.com/gocolly/colly"
)

var dolarPy = &Spider{
	Name:     "dolarpy",
	Selector: ".blue2",
	URL:      "https://www.dolarpy.com.br/",
	GetValue: func(e *colly.HTMLElement) string {
		return e.Text
	},
}
