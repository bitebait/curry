package spiders

import (
	"github.com/gocolly/colly"
)

var megaEletro = &Spider{
	Name:     "megaeletro",
	Selector: "div.mr-7:nth-child(3) > p:nth-child(2)",
	URL:      "https://www.megaeletro.com.py/br",
	GetValue: func(e *colly.HTMLElement) string {
		return e.Text
	},
}
