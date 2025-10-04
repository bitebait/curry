package spiders

func init() {
	NewSpider(
		"atacadocollections",
		"div.currency-results > span.text",
		"https://www.atacadocollections.com/",
	)
}
