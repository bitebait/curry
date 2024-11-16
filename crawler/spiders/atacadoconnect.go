package spiders

func init() {
	NewSpider(
		"atacadoconnect",
		"#j_idt24 > span",
		"https://www.atacadoconnect.com/",
	)
}
