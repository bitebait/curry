package spiders

func init() {
	NewSpider(
		"megaeletro",
		"div.mr-7:nth-child(3) > p:nth-child(2)",
		"https://www.megaeletro.com.py/br",
	)
}
