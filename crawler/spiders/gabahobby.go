package spiders

import (
	"github.com/gocolly/colly"
)

var gabbaHobby = &Spider{
	Name:     "gabahobby",
	Selector: ".special > ul:nth-child(1) > li:nth-child(3) > span:nth-child(2)",
	URL:      "https://www.gabahobby.com/",
	GetValue: func(e *colly.HTMLElement) string {
		return e.Text
	},
}
