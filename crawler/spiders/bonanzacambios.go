package spiders

import (
	"github.com/gocolly/colly"
)

var bonanzaCambios = &Spider{
	Name:     "bonanzacambios",
	Selector: "div.content-inner:nth-child(1) > div:nth-child(2) > div:nth-child(2) > section:nth-child(1) > div:nth-child(1) > div:nth-child(1) > div:nth-child(1) > div:nth-child(1) > table:nth-child(1) > tbody:nth-child(2) > tr:nth-child(3) > td:nth-child(5)",
	URL:      "https://www.bonanzacambios.com.py/",
	GetValue: func(e *colly.HTMLElement) string {
		return e.Text
	},
}
