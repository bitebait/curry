package spiders

func init() {
	NewSpider(
		"gabahobby",
		".special > ul:nth-child(1) > li:nth-child(3) > span:nth-child(2)",
		"https://www.gabahobby.com/",
	)
}
