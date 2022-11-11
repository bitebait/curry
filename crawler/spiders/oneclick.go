package spiders

import (
	"github.com/gocolly/colly"
)

var OneClick = &Spider{
	Name:     "oneclick",
	Selector: ".login-box > center:nth-child(1) > h5:nth-child(1)",
	URL:      "https://oneclick.com.py/",
	GetValue: func(e *colly.HTMLElement) string {
		return e.Text
	},
}
