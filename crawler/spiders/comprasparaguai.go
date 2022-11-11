package spiders

import (
	"github.com/gocolly/colly"
)

var comprasParaguai = &Spider{
	Name:     "comprasparaguai",
	Selector: "body > div > div.container-header-bottom > div > div > div.dolar-cotacao > div.btn-toggle-calculator.btn-toggle-calculator-js.flex > div > span.txt-quotation.hidden-xs > strong",
	URL:      "https://www.comprasparaguai.com.br/",
	GetValue: func(e *colly.HTMLElement) string {
		return e.Text
	},
}
