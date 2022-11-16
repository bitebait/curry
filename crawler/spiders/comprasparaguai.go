package spiders

import (
	"github.com/gocolly/colly"
)

func init() {
	NewSpider(
		"comprasparaguai",
		"body > div > div.container-header-bottom > div > div > div.dolar-cotacao > div.btn-toggle-calculator.btn-toggle-calculator-js.flex > div > span.txt-quotation.hidden-xs > strong",
		"https://www.comprasparaguai.com.br/",
		func(e *colly.HTMLElement) string {
			return e.Text
		},
	)
}
