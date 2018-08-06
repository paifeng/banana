package main

import (
	"testing"
	"os"
	"html/template"
	"log"
)

func Test_Template(t *testing.T) {
	tpl, err := template.ParseFiles("C:\\Users\\love5\\Desktop\\test.html")
	if err != nil {
		log.Fatal(err)
	}

	data := struct {
		Title string
	}{
		Title: "模板测试",
	}
	err = tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
}
