package spiders

import (
	"strings"

	"github.com/gocolly/colly"
)

var AudiumEletronics = &Spider{
	Name:     "audiumelectronics",
	Selector: "div.quotation:nth-child(3) > span:nth-child(1)",
	URL:      "https://www.audiumelectronics.com/home",
	GetValue: func(e *colly.HTMLElement) string {
		data := strings.TrimSpace(e.Text)
		return data
	},
}
