package spiders

import (
	"strings"

	"github.com/gocolly/colly"
)

var CellShop = &Spider{
	Name:     "cellshop",
	Selector: ".showCur",
	URL:      "https://www.cellshop.com/br/",
	GetValue: func(e *colly.HTMLElement) string {
		return strings.Split(e.Text, "Câmbio")[1]
	},
}
