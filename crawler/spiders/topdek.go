package spiders

import (
	"strings"

	"github.com/gocolly/colly"
)

var topDek = &Spider{
	Name:     "topdek",
	Selector: "body > div:nth-child(8) > div:nth-child(1) > p:nth-child(1)",
	URL:      "https://www.topdek.com.br/br",
	GetValue: func(e *colly.HTMLElement) string {
		return strings.Split(e.Text, " ")[3]
	},
}
