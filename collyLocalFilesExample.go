package main

import (
  "fmt"
  "net/http"
  "strings"
  "github.com/gocolly/colly"
)

func main() {
    t := &http.Transport{}
    t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))

    cFile := colly.NewCollector()
    cFile.WithTransport(t)

    cFile.OnHTML(main_tag, func(e *colly.HTMLElement) {
        // DO STUFF
    })

    fmt.Println("Visiting local file")
    cFile.Visit("file:" + "YOUR_PATH_TO_LOCAL_HTML_GILE")
}
