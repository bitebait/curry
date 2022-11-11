package spiders

import (
	"strings"

	"github.com/gocolly/colly"
)

var AtacadoCollections = &Spider{
	Name:     "atacadocollections",
	Selector: ".price",
	URL:      "https://www.atacadocollections.com/",
	GetValue: func(e *colly.HTMLElement) string {
		data := strings.TrimSpace(e.Text)
		return data
	},
}
