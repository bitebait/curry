package spiders

import (
	"strings"

	"github.com/gocolly/colly"
)

var TopDek = &Spider{
	Name:     "topdek",
	Selector: "body > div:nth-child(8) > div:nth-child(1) > p:nth-child(1)",
	URL:      "https://www.topdek.com.br/br",
	GetValue: func(e *colly.HTMLElement) string {
		data := strings.TrimSpace(strings.Split(e.Text, " ")[3])
		return data
	},
}
