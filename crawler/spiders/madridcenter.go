package spiders

import (
	"github.com/gocolly/colly"
)

var MadridCenter = &Spider{
	Name:     "madridcenter",
	Selector: "body > header > div > div > div > div.item.top-quotes > div.col-12.text3 > span > span:nth-child(1) > strong",
	URL:      "https://www.madridcenter.com/",
	GetValue: func(e *colly.HTMLElement) string {
		return e.Text
	},
}
