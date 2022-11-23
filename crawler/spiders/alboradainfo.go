package spiders

func init() {
	NewSpider(
		"alboradainfo",
		".quotation > span:nth-child(2)",
		"https://www.alboradainfo.com/",
	)
}
