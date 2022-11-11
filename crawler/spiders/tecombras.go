package spiders

import (
	"strings"

	"github.com/gocolly/colly"
)

var TeCombras = &Spider{
	Name:     "tecombras",
	Selector: ".currency-selector > span:nth-child(1)",
	URL:      "https://www.tecombras.net/",
	GetValue: func(e *colly.HTMLElement) string {
		data := strings.TrimSpace(strings.Split(e.Text, "=")[1])
		return data
	},
}
