package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var data []string
	var mobiles []string

	csvfile, _ := os.Open("output.csv")

	r := csv.NewReader(csvfile)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if len(record) > 0 {
			t_mobile := strings.Split(record[0], "|")
			mobile := t_mobile[0]
			if !IsPresent(&mobiles, mobile) {
				data = append(data, record[0])
			}
		}
	}
	WriteLines(data, "clean.csv")
}

func IsPresent(slice *[]string, element string) bool {
	for _, ele := range *slice {
		if ele == element {
			return true
		}
	}
	*slice = append(*slice, element)
	return false
}

func WriteLines(slice []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range slice {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}
