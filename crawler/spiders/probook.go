package spiders

import (
	"github.com/gocolly/colly"
)

var proBook = &Spider{
	Name:     "probook",
	Selector: ".quotation > span:nth-child(1) > strong:nth-child(1)",
	URL:      "https://www.probook.com.py/",
	GetValue: func(e *colly.HTMLElement) string {
		return e.Text
	},
}
