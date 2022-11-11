package spiders

import (
	"github.com/gocolly/colly"
)

var AtacadoCollections = &Spider{
	Name:     "atacadocollections",
	Selector: ".price",
	URL:      "https://www.atacadocollections.com/",
	GetValue: func(e *colly.HTMLElement) string {
		return e.Text
	},
}
