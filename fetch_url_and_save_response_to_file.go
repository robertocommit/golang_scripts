package main

import (
	"fmt"
	"io/ioutil"
  "net/http"
  "os"
)

func main() {
  url := "INSERT_URL_HERE"
  res, err := http.Get(url)
  fmt.Println("Visiting", url)
  if err != nil {
      panic(err.Error())
  }
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
      panic(err.Error())
  }
  SaveResponseToFile(string(body))
}

func SaveResponseToFile(response string) {
  fmt.Println("Saving response to file...")
  dir, err := os.Getwd()
  if err != nil {
    panic(err.Error())
  }
  f, err := os.Create(dir + "/response.html")
  if err != nil {
    panic(err.Error())
  }
  defer f.Close()
  f.WriteString(response)
}
