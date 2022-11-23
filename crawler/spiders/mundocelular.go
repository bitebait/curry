package spiders

func init() {
	NewSpider(
		"mundodocelular",
		".heading > div:nth-child(1) > span:nth-child(4)",
		"https://www.mundodocelular.com/",
	)
}
