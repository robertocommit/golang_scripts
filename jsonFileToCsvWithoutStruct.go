package main

import (
	"encoding/json"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

func main() {

    input_file_name := "allData.json"
    output_file_name := "outputFileName.csv"
    separator := "|||"

	var output []string
	var keys []string

	fileIn, _ := ioutil.ReadFile(input_file_name)

	var results []map[string]interface{}

	json.Unmarshal([]byte(fileIn), &results)

	for _, elem := range results {
		infos := elem["Infos"]
		iter := reflect.ValueOf(infos).MapRange()
		for iter.Next() {
			k := iter.Key().Interface().(string)
			keys = AppendIfMissing(keys, k)
		}
	}

	output = append(output, strings.Join(keys, separator))

	number_rows := len(keys)

	for _, elem := range results {
		infos := elem["Infos"]
		iter := reflect.ValueOf(infos).MapRange()
		row := make([]string, number_rows)
		for iter.Next() {
			position := IndexOf(keys, iter.Key().Interface().(string))
			v := iter.Value().Interface().(string)
			row[position] = v
		}
		row_merge := strings.Join(row, "|||")
		output = append(output, row_merge)
	}

	WriteLines(output, output_file_name)
}

func AppendIfMissing(slice []string, element string) []string {
	for _, ele := range slice {
		if ele == element {
			return slice
		}
	}
	return append(slice, element)
}

func IndexOf(slice []string, element string) int {
	for k, v := range slice {
		if element == v {
			return k
		}
	}
	return -1 // not found.
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
