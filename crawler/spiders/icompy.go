package spiders

import (
	"strings"

	"github.com/gocolly/colly"
)

func init() {
	NewSpider(
		"icompy",
		"#header > div > div > div > section.elementor-element.elementor-element-cldqzgw.elementor-section-stretched.elementor-section-boxed.elementor-section-height-default.elementor-section-content-middle.sticky-inner.elementor-hidden-tablet.elementor-hidden-phone.elementor-section.elementor-top-section > div > div > div.elementor-element.elementor-element-ligyjdn.elementor-column.elementor-col-25.elementor-top-column > div > div > div > div > div > div.elementor-icon-box-content > div",
		"http://icompy.com/",
		func(e *colly.HTMLElement) string {
			return strings.Split(e.Text, "G$")[0]
		},
	)
}
