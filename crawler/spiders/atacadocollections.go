package spiders

func init() {
	NewSpider(
		"atacadocollections",
		"#header > div.header-extras > div > ul:nth-child(1) > li:nth-child(1) > span",
		"https://www.atacadocollections.com/",
	)
}
