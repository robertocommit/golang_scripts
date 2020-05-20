package main

import (
    "fmt"
)

type Result struct {
    Name    string
    Url     string
}

func main() {
    m := map[Result]int{
        Result{"Test_1", "https://www.gooo_1.com"}: 1,
        Result{"Test_2", "https://www.gooo_2.com"}: 2,
        Result{"Test_3", "https://www.gooo_3.com"}: 3,
        Result{"Test_4", "https://www.gooo_4.com"}: 4,
        Result{"Test_1", "https://www.gooo_1.com"}: 5,
    }
    for key, value := range m {
        fmt.Println("key:", key, "value:", value)
    }
}
