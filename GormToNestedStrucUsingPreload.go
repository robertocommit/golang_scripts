package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
    ID       int
    Username string
    Session []Session `gorm:"ForeignKey:Userid"`
}

type Session struct {
	ID     int
	Email  string
	Userid int
}

func main() {
	db, _ := gorm.Open("XXXX", "host=XXXXXX port=5432 user=XXXXX dbname=XXXXX password=XXXXX")
	var s []Session
	db.Preload("User").Find(&s)

	fmt.Println(s)
}
