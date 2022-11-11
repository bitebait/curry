package spiders

import (
	"strings"

	"github.com/gocolly/colly"
)

var VisaoVip = &Spider{
	Name:     "visaovip",
	Selector: "div.ui-panelgrid-cell:nth-child(1) > span:nth-child(1)",
	URL:      "http://visaovip.com/",
	GetValue: func(e *colly.HTMLElement) string {
		data := strings.TrimSpace(e.Text)
		return data
	},
}
