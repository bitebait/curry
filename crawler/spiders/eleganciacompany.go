package spiders

func init() {
	NewSpider(
		"eleganciacompany",
		"div.quotation span",
		"https://www.eleganciacompany.com/",
	)
}
