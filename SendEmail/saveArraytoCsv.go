package main

import (
    "os"
    "encoding/csv"
)

func main() {
    var data [][]string
    row := []string{"a", "b", "c", "d"}
    data = append(data, row)
    scrape("example.csv", data)
}

func SaveRowToCsv(filename string, data [][]string) {
    f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
    if err != nil {
        panic(err.Error())
    }
    defer f.Close()
    w := csv.NewWriter(f)
    w.Comma = ';'
    w.WriteAll(data)
    if err := w.Error(); err != nil {
        panic(err.Error())
    }
}