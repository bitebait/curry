package spiders

import (
	"github.com/gocolly/colly"
)

var tcheLoco = &Spider{
	Name:     "tcheloco",
	Selector: ".nohover",
	URL:      "https://www.tcheloco.com.py/br/",
	GetValue: func(e *colly.HTMLElement) string {
		return e.Text
	},
}
