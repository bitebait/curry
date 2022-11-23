package spiders

func init() {
	NewSpider(
		"atlanticoshop",
		"div.bandera:nth-child(3) > small:nth-child(2)",
		"https://www.atlanticoshop.com.br/",
	)
}
