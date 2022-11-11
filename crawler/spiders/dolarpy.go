package spiders

import (
	"strings"

	"github.com/gocolly/colly"
)

var DolarPy = &Spider{
	Name:     "dolarpy",
	Selector: ".blue2",
	URL:      "https://www.dolarpy.com.br/",
	GetValue: func(e *colly.HTMLElement) string {
		data := strings.TrimSpace(e.Text)
		return data
	},
}
