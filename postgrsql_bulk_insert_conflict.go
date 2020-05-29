package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
	"strconv"
    "time"
	"strings"
)

func InsertResults(Db *sql.DB, results []Result) {
	valueStrings := []string{}
	valueArgs := []interface{}{}
	for i, elem := range results {
        str1 := "$" + strconv.Itoa(1+i*6) + ","
        str2 := "$" + strconv.Itoa(2+i*6) + ","
        str3 := "$" + strconv.Itoa(3+i*6) + ","
        str4 := "$" + strconv.Itoa(4+i*6) + ","
        str5 := "$" + strconv.Itoa(5+i*6) + ","
        str6 := "$" + strconv.Itoa(6+i*6)
        str_n := "(" + str1 + str2 + str3 + str4 + str5 + str6 + ")"
        valueStrings = append(valueStrings, str_n)
		valueArgs = append(valueArgs, elem.Value_1)
		valueArgs = append(valueArgs, elem.Value_2)
		valueArgs = append(valueArgs, elem.Value_3)
		valueArgs = append(valueArgs, elem.Value_4)
		valueArgs = append(valueArgs, elem.Value_5)
		valueArgs = append(valueArgs, elem.Value_6)
	}
    smt := `INSERT INTO results (
                value_1, value_2, value_3, value_4, value_5, value_6)
            VALUES %s ON CONFLICT (value_1) DO UPDATE
            SET value_2 = EXCLUDED.value_2,
                value_3 = EXCLUDED.value_3,
                value_4 = EXCLUDED.value_4`
    smt = fmt.Sprintf(smt, strings.Join(valueStrings, ","))
    _, err := Db.Exec(smt, valueArgs...)
	if err != nil {
		panic(err.Error())
	}
}
