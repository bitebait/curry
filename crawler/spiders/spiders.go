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

var spiderRunners []Runnable

type Runnable interface {
	RunSpider(channel chan models.Store)
}

func GetAllSpiders() []Runnable {
	return spiderRunners
}

type Spider struct {
	Name     string
	Selector string
	GetValue func(e *colly.HTMLElement) string
	URL      string
}

func NewSpider(name, selector, url string, getValue ...func(e *colly.HTMLElement) string) {
	s := Spider{
		Name:     name,
		Selector: selector,
		URL:      url,
	}
	if len(getValue) == 0 {
		s.GetValue = defaultGetValue
	} else {
		s.GetValue = getValue[0]
	}
	spiderRunners = append(spiderRunners, s)
}

func defaultGetValue(e *colly.HTMLElement) string {
	return e.Text
}

func (s Spider) RunSpider(channel chan models.Store) {
	c := createCollector()
	configureCollector(s, c, channel)

	err := c.Visit(s.URL)
	if err != nil {
		log.Fatalf("FATAL: Failed to visit URL: %s\n", s.URL)
	}

	c.Wait()
}

func createCollector() *colly.Collector {
	collector := colly.NewCollector(colly.Async(true))
	extensions.RandomUserAgent(collector)
	extensions.Referer(collector)
	collector.WithTransport(&http.Transport{
		DisableKeepAlives:     true,
		ResponseHeaderTimeout: time.Duration(config.GetConfig.Crawler.ResponseHeaderTimeout) * time.Second,
	})
	collector.SetRequestTimeout(time.Duration(config.GetConfig.Crawler.ClientTimeout) * time.Second)
	collector.Limit(&colly.LimitRule{Parallelism: 6})
	return collector
}

func configureCollector(s Spider, c *colly.Collector, channel chan models.Store) {
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
}
