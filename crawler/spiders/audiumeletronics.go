package spiders

import (
	"github.com/gocolly/colly"
)

var audiumEletronics = &Spider{
	Name:     "audiumelectronics",
	Selector: "div.quotation:nth-child(3) > span:nth-child(1)",
	URL:      "https://www.audiumelectronics.com/home",
	GetValue: func(e *colly.HTMLElement) string {
		return e.Text
	},
}
