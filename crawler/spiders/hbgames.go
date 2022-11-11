package spiders

import (
	"github.com/gocolly/colly"
)

var HBGames = &Spider{
	Name:     "hbgames",
	Selector: "body > div.main.document > header.header.site.show-for-medium > div.bar.top > div > div.first.row.grid-x > div > div > div.quotation.cell.small-30.medium-25.large-12 > div.value > span:nth-child(2)",
	URL:      "http://www.hbgamespy.com/",
	GetValue: func(e *colly.HTMLElement) string {
		return e.Text
	},
}
