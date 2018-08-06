package main

import (
	"testing"
	"github.com/armon/go-radix"
	"fmt"
)

func sayWord(word string) {
	fmt.Println(word)
}

func Test_Radix(t *testing.T) {
	r := radix.New()
	r.Insert("/user", sayWord)
	ret, ok := r.Get("/user")
	if ok {
		f := ret.(func(string))
		f("OKOK")
	}
	r.ToMap()
}
