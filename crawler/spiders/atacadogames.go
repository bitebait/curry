package spiders

import (
	"strings"

	"github.com/gocolly/colly"
)

var atacadoGames = &Spider{
	Name:     "atacadogames",
	Selector: "div.header-extra div.wrap.is-centralized ul.grid li.grid-item span.text.has-icon",
	URL:      "https://www.atacadogames.com/",
	GetValue: func(e *colly.HTMLElement) string {
		return strings.Split(e.Text, "Cotação")[1]
	},
}
