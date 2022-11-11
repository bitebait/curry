package spiders

import (
	"github.com/gocolly/colly"
)

var LGImportados = &Spider{
	Name:     "lgimportados",
	Selector: ".quotation > span:nth-child(1) > strong:nth-child(1)",
	URL:      "https://www.lgimportados.com/",
	GetValue: func(e *colly.HTMLElement) string {
		return e.Text
	},
}
