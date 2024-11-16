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

// Global variable to store a list of Runnable spiders
var spiderRunners []Runnable

// Runnable interface
type Runnable interface {
	RunSpider(channel chan models.Store)
}

// Function to retrieve all spiders
func GetAllSpiders() []Runnable {
	return spiderRunners
}

// Spider struct definition
type Spider struct {
	Name     string
	Selector string
	GetValue func(e *colly.HTMLElement) string
	URL      string
}

// Function to create a new spider
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

// Default function to get value from HTML element
func defaultGetValue(e *colly.HTMLElement) string {
	return e.Text
}

// Function to run the spider
func (s Spider) RunSpider(channel chan models.Store) {
	c := createCollector()
	configureCollector(s, c, channel)

	err := c.Visit(s.URL)
	if err != nil {
		log.Fatalf("FATAL: Failed to visit URL: %s\n", s.URL)
	}

	c.Wait()
}

// Function to create a new collector with necessary configurations
func createCollector() *colly.Collector {
	collector := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(1),
	)
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

// Global map to track seen URLs
var seenURLs = make(map[string]bool)

// Function to configure the collector with necessary callbacks
func configureCollector(s Spider, c *colly.Collector, channel chan models.Store) {
	c.OnHTML(s.Selector, func(e *colly.HTMLElement) {
		value := s.GetValue(e)

		// Check for duplicate URLs before sending to channel
		if seenURLs[s.URL] {
			return
		}

		seenURLs[s.URL] = true

		store := &models.Store{
			Name:     s.Name,
			Currency: config.GetConfig.Currency.Currency,
			Value:    helpers.FormatCurrency(value),
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
