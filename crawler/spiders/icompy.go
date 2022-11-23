package spiders

func init() {
	NewSpider(
		"icompy",
		"div.pt_menu:nth-child(2) > div:nth-child(1) > a:nth-child(1) > span:nth-child(1)",
		"http://icompy.com/",
	)
}
