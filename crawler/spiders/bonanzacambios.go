package spiders

func init() {
	NewSpider(
		"bonanzacambios",
		"div.content-inner:nth-child(1) > div:nth-child(2) > div:nth-child(2) > section:nth-child(1) > div:nth-child(1) > div:nth-child(1) > div:nth-child(1) > div:nth-child(1) > table:nth-child(1) > tbody:nth-child(2) > tr:nth-child(3) > td:nth-child(5)",
		"https://www.bonanzacambios.com.py/",
	)
}
