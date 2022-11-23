package spiders

func init() {
	NewSpider(
		"lgimportados",
		".quotation > span:nth-child(1) > strong:nth-child(1)",
		"https://www.lgimportados.com/",
	)
}
