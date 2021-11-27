package config

import (
	"curry/api/models"

	"github.com/gocolly/colly"
)

type Spider struct {
	Name     string
	Channel  chan models.Store
	Selector string
	GetValue func(e *colly.HTMLElement) string
	URL      string
}
