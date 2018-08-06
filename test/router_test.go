package main

import (
	"github.com/armon/go-radix"
	"testing"
	"fmt"
)

type Router struct {
	tree *radix.Tree
}

type Handle func(string)

func New() *Router {
	return &Router{radix.New()}
}

func (router *Router) Get(path string) (Handle, bool) {
	f, ok := router.tree.Get(path)
	return f.(Handle), ok
}

func (router *Router) Insert(path string, handle Handle) bool {
	_, ok := router.tree.Insert(path, handle)
	return ok
}

func helloWorld(word string) {
	fmt.Println(word)
}

func Test_Router(t *testing.T) {
	router := New()
	router.Insert("/user", helloWorld)
	f, ok := router.Get("/user")
	if ok {
		f("Hello World")
	}
}
