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
		data := strings.TrimSpace(strings.Split(e.Text, "Câmbio")[1])
		return data
	},
}
