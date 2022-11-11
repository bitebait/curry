package spiders

import (
	"strings"

	"github.com/gocolly/colly"
)

var PioneerInter = &Spider{
	Name:     "pioneerinter",
	Selector: ".currency-selector > span:nth-child(1)",
	URL:      "https://www.pioneerinter.com/",
	GetValue: func(e *colly.HTMLElement) string {
		data := strings.TrimSpace(strings.Split(e.Text, "=")[1])
		return data
	},
}
