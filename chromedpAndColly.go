package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly"
	. "github.com/logrusorgru/aurora"
)

func main() {

	start_url := "YOUR_URL"
	file_name := "YOUR_FILE_NAME"

	// CHROMEDP GET HTML CODE
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	var initialResponse string
	if err := chromedp.Run(ctx,
		chromedp.Navigate(start_url),
		chromedp.WaitVisible("YOUR_FIRST_TAG"),
		chromedp.OuterHTML("html", &initialResponse),
	); err != nil {
		panic(err.Error())
	}
	SaveResponseToFileWithFileName(initialResponse, file_name)

	// COLLY SCRAPE INFORMATION FROM HTML
	c := colly.NewCollector()
	c.OnHTML("YOUR_SECOND_TAG", func(e *colly.HTMLElement) {
		fmt.Println(e)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println(Gray(8-1, "Visiting"), Gray(8-1, r.URL.String()))
	})
	c.OnScraped(func(r *colly.Response) {
		RemoveFileWithFileName(file_name)
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println(Red("Request URL:"), Red(r.Request.URL))
	})

	// COLLY OPEN HTML FILE
	t := &http.Transport{}
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))
	dir, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	c.WithTransport(t)
	c.Visit("file:" + dir + "/" + file_name)
	return
}

func SaveResponseToFileWithFileName(response string, filename string) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	f, err := os.Create(dir + "/" + filename)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()
	f.WriteString(response)
}

func RemoveFileWithFileName(filename string) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	err = os.Remove(dir + "/" + filename)
	if err != nil {
		panic(err.Error())
	}
}
