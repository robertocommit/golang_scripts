package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

type Job struct {
	Title string
	Url   string
}

type Scraper struct {
	Url     string
	Tag     string
	Attr    string
	Value   string
	TextTag string
	UrlTag  string
}

func (scraper *Scraper) scrape() (jobs []Job, err error) {

	jobList := colly.NewCollector()

	jobList.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	jobList.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})
	jobList.OnError(func(_ *colly.Response, err error) {
		return
	})
	jobList.OnHTML(scraper.Tag, func(e *colly.HTMLElement) {
		if strings.Contains(e.Attr(scraper.Attr), scraper.Value) {
			title := e.ChildText(scraper.TextTag)
			url := e.Attr("href")
			if url == "" {
				url = e.ChildAttr(scraper.UrlTag, "href")
			}
			temp_job := Job{title, url}
			jobs = append(jobs, temp_job)
		}
	})

	jobList.Visit(scraper.Url)
	return
}

func main() {
	runner := Scraper{"https://www.kununu.com/at/kununu/jobs", "div", "class", "company-profile-job-item", "a", "a"}
	// runner := Scraper{"https://careers.artnight.com/", "a", "class", "col-md-6", "h5", "title"}
	// runner := Scraper{"https://www.papyrus.de/karriere", "div", "class", "et_pb_with_border", "h4", "a"}
	jobs, err := runner.scrape()
	if err == nil {
		fmt.Println(jobs)
	} else {
		fmt.Println("Program terminated with err: ", err)
	}
}
