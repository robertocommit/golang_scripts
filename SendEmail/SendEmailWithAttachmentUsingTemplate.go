package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

type info struct {
	Name string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
	password := os.Getenv("PASSWORD")

	t := template.New("template.html")

	t, err = t.ParseFiles("template.html")
	if err != nil {
		fmt.Println(err)
	}

	i := info{"Roberto"}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, i); err != nil {
		fmt.Println(err)
	}

	result := tpl.String()

	m := gomail.NewMessage()
	m.SetHeader("From", "robimalco@gmail.com")
	m.SetHeader("To", "robimalco@gmail.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", result)
	m.Attach("test.jpg")

	d := gomail.NewDialer("smtp.gmail.com", 587, "robimalco@gmail.com", password)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
