package spiders

func init() {
	NewSpider(
		"visaovip",
		"div.ui-panelgrid-cell:nth-child(1) > span:nth-child(1)",
		"http://visaovip.com/",
	)
}
