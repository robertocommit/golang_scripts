package main

import (
	"encoding/json"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
    "strings"
    "regexp"
)

func main() {

    input_file_name := "allData.json"
    output_file_name := "outputFileName.csv"
    separator := "|"

    // Prepare regex to clean each row fron newlines
    reg, err := regexp.Compile(`(?m)^\n*$`)
    if err != nil {
        fmt.Println(err)
    }

	var output []string
    
	fileIn, _ := ioutil.ReadFile(input_file_name)
    
	var results []map[string]interface{}
    
	json.Unmarshal([]byte(fileIn), &results)
    
    // Get columns' headers
	var keys []string
	for _, elem := range results {
		infos := elem["Infos"]
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
		infos := elem["Infos"]
		iter := reflect.ValueOf(infos).MapRange()
		row := make([]string, number_rows)
		for iter.Next() {
            key := iter.Key().Interface().(string)
			position := IndexOf(keys, key)
            v := iter.Value().Interface().(string)
            processedV := reg.ReplaceAllString(v, "")
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
