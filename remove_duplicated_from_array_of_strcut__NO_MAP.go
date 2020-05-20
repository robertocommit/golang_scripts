package main

import (
    "fmt"
)

type Result struct {
    Name    string
    Url     string
}

func main() {
    result_1 := Result{"Test_1", "https://www.gooo_1.com"}
    result_2 := Result{"Test_2", "https://www.gooo_2.com"}
    result_3 := Result{"Test_3", "https://www.gooo_3.com"}
    result_4 := Result{"Test_4", "https://www.gooo_4.com"}
    result_5 := Result{"Test_1", "https://www.gooo_1.com"}

    var results = []Result {result_1, result_2, result_3, result_4, result_5}
    var unique_results = unique(results)

    fmt.Println(unique_results)
}

func unique(duplicatedResults []Result) []Result {
    var uniqueResults []Result

    for _, v := range duplicatedResults {
        skip := false
        for _, u := range uniqueResults {
            if v.Name == u.Name && v.Url == u.Url {
                skip = true
                break
            }
        }
        if !skip {
            uniqueResults = append(uniqueResults, v)
        }
    }
    return uniqueResults
}
