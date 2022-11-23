package spiders

func init() {
	NewSpider(
		"probook",
		".quotation > span:nth-child(1) > strong:nth-child(1)",
		"https://www.probook.com.py/",
	)
}
