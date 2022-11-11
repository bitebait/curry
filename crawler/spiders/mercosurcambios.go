package spiders

import (
	"github.com/gocolly/colly"
)

var mercosurCambios = &Spider{
	Name:     "mercosurcambios",
	Selector: "div.tab_panel:nth-child(2) > div:nth-child(1) > div:nth-child(1) > div:nth-child(2) > div:nth-child(1) > div:nth-child(1) > div:nth-child(2) > table:nth-child(1) > tbody:nth-child(2) > tr:nth-child(3) > th:nth-child(5)",
	URL:      "https://site.mercosurcambios.com/",
	GetValue: func(e *colly.HTMLElement) string {
		return e.Text
	},
}
