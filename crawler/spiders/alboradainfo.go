package spiders

import "github.com/gocolly/colly"

func init() {
	NewSpider(
		"alboradainfo",
		".quotation > span:nth-child(2)",
		"https://www.alboradainfo.com/",
		func(e *colly.HTMLElement) string {
			return e.Text
		},
	)
}
