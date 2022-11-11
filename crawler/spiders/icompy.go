package spiders

import (
	"github.com/gocolly/colly"
)

var Icompy = &Spider{
	Name:     "icompy",
	Selector: "div.pt_menu:nth-child(2) > div:nth-child(1) > a:nth-child(1) > span:nth-child(1)",
	URL:      "http://icompy.com/",
	GetValue: func(e *colly.HTMLElement) string {
		return e.Text
	},
}
