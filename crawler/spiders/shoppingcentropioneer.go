package spiders

func init() {
	NewSpider(
		"shoppingcentropioneer",
		"ul.quotation:nth-child(2) > li:nth-child(1) > span:nth-child(2)",
		"https://shoppingcentropioneer.com/",
	)
}
