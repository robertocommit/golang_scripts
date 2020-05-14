package main

import (
    "fmt"
    "github.com/gocolly/colly"
)

func main() {

	jobList := colly.NewCollector(
        colly.MaxDepth(2),
        colly.AllowedDomains("boards.greenhouse.io"),
    )
    
    jobDescriptions := jobList.Clone()

    jobList.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL)
    })
    jobList.OnScraped(func(r *colly.Response) {
        fmt.Println("Finished", r.Request.URL)
    })

    jobList.OnHTML("a", func(e *colly.HTMLElement) {
        sub_url := e.Request.AbsoluteURL(e.Attr("href"))
        jobDescriptions.Visit(sub_url)
    })

    jobDescriptions.OnHTML("h1", func(e *colly.HTMLElement) {
        if e.Attr("class") == "app-title" {
            fmt.Println("Job title --> ", e.Text)
        }
    })

    jobList.Visit("https://boards.greenhouse.io/yoursuper?")
}
