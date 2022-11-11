package spiders

import (
	"strings"

	"github.com/gocolly/colly"
)

var VictoriaStore = &Spider{
	Name:     "victoriastore",
	Selector: ".left-content > div:nth-child(1) > ul:nth-child(1) > li:nth-child(2) > div:nth-child(1) > span:nth-child(1)",
	URL:      "https://www.victoriastore.com.br/",
	GetValue: func(e *colly.HTMLElement) string {
		data := strings.TrimSpace(strings.Split(e.Text, "=")[1])
		return data
	},
}
