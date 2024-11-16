package spiders

func init() {
	NewSpider(
		"topdek",
		"#navbarDropdown > strong:nth-child(1)",
		"https://www.topdek.com.br/br",
	)
}
