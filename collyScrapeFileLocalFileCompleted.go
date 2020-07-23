package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	type Job struct {
		Url        string
		Title      string
		Department string
	}

	c.OnHTML(".work-openings", func(e *colly.HTMLElement) {
		e.ForEach(".list-box-item", func(_ int, el *colly.HTMLElement) {
			result_title := el.ChildText("a")
			result_url := el.ChildAttr("a", "href")
			fmt.Println(result_title, result_url)
		})
	})

	t := &http.Transport{}
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))
	c.WithTransport(t)
	dir, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	c.Visit("file:" + dir + "/response.html")
}
