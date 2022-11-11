package spiders

import (
	"github.com/gocolly/colly"
)

var MundoCelular = &Spider{
	Name:     "mundodocelular",
	Selector: ".heading > div:nth-child(1) > span:nth-child(4)",
	URL:      "https://www.mundodocelular.com/",
	GetValue: func(e *colly.HTMLElement) string {
		return e.Text
	},
}
