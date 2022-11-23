package spiders

func init() {
	NewSpider(
		"infinitysport",
		"body > header > div > div > div > div.quotation > div.info > span",
		"https://www.infinitysport.com.py/", 
	)
}
