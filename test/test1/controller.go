package main

import (
	"github.com/banana/ba"
	"fmt"
	"reflect"
	"strings"
)

type User struct {
	login int `method:"get" router:"/user/login"`
	new   int `method:"get" router:"/user/new"`
	index int `method:"get" router:"/user/index"`
}

func (u *User) LOGIN(req ba.Req, resp ba.Resp) {
	fmt.Print("LOGIN")
}

func (u *User) NEW(req ba.Req, resp ba.Resp) {
	fmt.Print("NEW")
}

func (u *User) INDEX(req ba.Req, resp ba.Resp) {
	fmt.Print("INDEX")
}

func INDEX(req ba.Req, resp ba.Resp) {
	fmt.Print("INDEX")
}

func main() {
	u := &User{}
	tof := reflect.TypeOf(u)
	tofEle := tof.Elem()
	numOfFields := tofEle.NumField()
	for i := 0; i < numOfFields; i++ {
		name := tofEle.Field(i).Name
		method := tofEle.Field(i).Tag.Get("method")
		router := tofEle.Field(i).Tag.Get("router")
		fun, _ := tof.MethodByName(strings.ToUpper(name))
		f := fun.Func.Interface().(func(*User, ba.Req, ba.Resp))
		f(u, nil, nil)
		fmt.Print(" " + method)
		fmt.Print(" " + router + "\n")
	}
}
