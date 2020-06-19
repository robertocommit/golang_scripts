package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	. "github.com/logrusorgru/aurora"
)

func main() {
	start := time.Now()

	m_start_url := "https://careers.microsoft.com/us/en/search-results?s=1&from="

	c := colly.NewCollector()

	q, _ := queue.New(
		5, // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
	)

	c.OnHTML("html", func(e *colly.HTMLElement) {
		temp_results, _ := e.DOM.Find("link[rel=canonical]").Attr("href")
		fmt.Println("\t", Green(temp_results))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	for i := 1; i < 10; i++ {
		q.AddURL(m_start_url + strconv.Itoa(i*10))
	}
	q.Run(c)

	elapsed := time.Since(start)
	fmt.Println(Blue("Took %s"), Blue(elapsed))
}
