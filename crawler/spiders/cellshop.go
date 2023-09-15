package spiders

func init() {
	NewSpider(
		"cellshop",
		"div.currency-exchange-rates > span.rate",
		"https://www.cellshop.com/br/",
	)
}
