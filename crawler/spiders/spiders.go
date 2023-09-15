package spiders

import (
	"log"
	"net/http"
	"time"

	"github.com/bitebait/curry/api/models"
	"github.com/bitebait/curry/config"
	"github.com/bitebait/curry/helpers"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

var spiders []Runnable

type Runnable interface {
	RunSpider(channel chan models.Store)
}

func GetAllSpiders() []Runnable {
	return spiders
}

type spider struct {
	Name     string
	Selector string
	GetValue func(e *colly.HTMLElement) string
	URL      string
}

func NewSpider(name, selector, url string, getValue ...func(e *colly.HTMLElement) string) {
	s := spider{
		Name:     name,
		Selector: selector,
		URL:      url,
	}
	if len(getValue) <= 0 {
		s.GetValue = func(e *colly.HTMLElement) string {
			return e.Text
		}
	} else {
		s.GetValue = getValue[0]
	}
	spiders = append(spiders, s)
}

func (s spider) RunSpider(channel chan models.Store) {
	c := colly.NewCollector(colly.Async(true))

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	c.WithTransport(&http.Transport{
		DisableKeepAlives:     true,
		ResponseHeaderTimeout: time.Duration(config.GetConfig.Crawler.ResponseHeaderTimeout) * time.Second,
	})

	c.SetRequestTimeout(time.Duration(config.GetConfig.Crawler.ClientTimeout) * time.Second)

	c.Limit(&colly.LimitRule{Parallelism: 6})

	c.OnHTML(s.Selector, func(e *colly.HTMLElement) {
		store := &models.Store{
			Name:     s.Name,
			Currency: config.GetConfig.Currency.Currency,
			Value:    helpers.FormatCurrency(s.GetValue(e)),
			URL:      s.URL,
		}
		channel <- *store
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Printf("FATAL: Something went wrong (%s): %s", r.Request.URL, err)
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("INFO: Visiting", r.URL)
	})

	err := c.Visit(s.URL)
	if err != nil {
		log.Fatalf("FATAL: Failed to visit URL: %s\n", s.URL)
	}

	c.Wait()
}
