package spiders

import (
	"strings"

	"github.com/gocolly/colly"
)

var casaAmericana = &Spider{
	Name:     "casaamericana",
	Selector: "span.justify-center:nth-child(2)",
	URL:      "https://www.casaamericana.com.py/",
	GetValue: func(e *colly.HTMLElement) string {
		return strings.Split(e.Text, "R$")[1]
	},
}
