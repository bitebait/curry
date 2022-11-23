package spiders

func init() {
	NewSpider(
		"comprasparaguai",
		"body > div > div.container-header-bottom > div > div > div.dolar-cotacao > div.btn-toggle-calculator.btn-toggle-calculator-js.flex > div > span.txt-quotation.hidden-xs > strong",
		"https://www.comprasparaguai.com.br/",
	)
}
