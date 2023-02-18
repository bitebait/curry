package spiders

func init() {
	NewSpider(
		"hbgames",
		".arial > div:nth-child(2) > b:nth-child(1)",
		"http://www.hbgamespy.com/",
	)
}
