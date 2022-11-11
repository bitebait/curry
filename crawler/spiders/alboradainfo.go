package spiders

import (
	"strings"

	"github.com/gocolly/colly"
)

var AlboradaInfo = &Spider{
	Name:     "alboradainfo",
	Selector: ".quotation > span:nth-child(2)",
	URL:      "https://www.alboradainfo.com/",
	GetValue: func(e *colly.HTMLElement) string {
		data := strings.TrimSpace(e.Text)
		return data
	},
}
