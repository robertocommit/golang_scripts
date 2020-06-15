// This script works (ATM) only with fully stringly Structs

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

func main() {
	input_file_name := "output.json"
	output_file_name := "output.csv"
	separator := "|"

	var output []string

	fileIn, _ := ioutil.ReadFile(input_file_name)

	var results []map[string]interface{}

	json.Unmarshal([]byte(fileIn), &results)

	// Get columns' headers
	var keys []string
	for _, elem := range results {
		infos := elem
		iter := reflect.ValueOf(infos).MapRange()
		for iter.Next() {
			k := iter.Key().Interface().(string)
			keys = AppendIfMissing(keys, k)
		}
	}

	// Append headers' to output
	output = append(output, strings.Join(keys, separator))

	number_rows := len(keys)

	// For each row assign a the value to the column it belongs
	for _, elem := range results {
		infos := elem
		iter := reflect.ValueOf(infos).MapRange()
		row := make([]string, number_rows)
		for iter.Next() {
			key := iter.Key().Interface().(string)
			position := IndexOf(keys, key)
			v := iter.Value().Interface().(string)
			processedV := strings.ReplaceAll(v, "\n", " ")
			row[position] = processedV
		}
		row_merge := strings.Join(row, separator)
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
	return -1
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
