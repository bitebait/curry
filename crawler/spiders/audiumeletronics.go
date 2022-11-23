package spiders

func init() {
	NewSpider(
		"audiumelectronics",
		"div.quotation:nth-child(3) > span:nth-child(1)",
		"https://www.audiumelectronics.com/home",
	)
}
