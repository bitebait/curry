package spiders

import (
	"github.com/gocolly/colly"
)

var AlboradaInfo = &Spider{
	Name:     "alboradainfo",
	Selector: ".quotation > span:nth-child(2)",
	URL:      "https://www.alboradainfo.com/",
	GetValue: func(e *colly.HTMLElement) string {
		return e.Text
	},
}
