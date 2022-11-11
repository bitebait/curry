package spiders

import (
	"github.com/gocolly/colly"
)

var ShoppingCentro = &Spider{
	Name:     "shoppingcentropioneer",
	Selector: "ul.quotation:nth-child(2) > li:nth-child(1) > span:nth-child(2)",
	URL:      "https://shoppingcentropioneer.com/",
	GetValue: func(e *colly.HTMLElement) string {
		return e.Text
	},
}
