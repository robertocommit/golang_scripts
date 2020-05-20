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

    results_unique := unique(results)
    
    fmt.Println(results_unique)
}

func unique(result []Result) []Result {
    var unique []Result
    type key struct{ Name, Url string }
    m := make(map[key]int)
    for _, v := range result {
        k := key{v.Name, v.Url}
        if i, ok := m[k]; ok {
            unique[i] = v
        } else {
            m[k] = len(unique)
            unique = append(unique, v)
        }
    }
    return unique
}
