package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	ID       int
	Username string
}

type Session struct {
	ID     int
	Email  string
	Userid int
	User   User `gorm:"ForeignKey:Userid;AssociationForeignKey:ID"`
}

func main() {
	db, _ := gorm.Open("XXXX", "host=XXXXXX port=5432 user=XXXXX dbname=XXXXX password=XXXXX")
	var s []Session
	db.Preload("User").Find(&s)

	fmt.Println(s)
}
