package spiders

func init() {
	NewSpider(
		"hbgames",
		"body > div.main.document > header.header.site.show-for-medium > div.bar.top > div > div.first.row.grid-x > div > div > div.quotation.cell.small-30.medium-25.large-12 > div.value > span:nth-child(2)",
		"http://www.hbgamespy.com/",
	)
}
