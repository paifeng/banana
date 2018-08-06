package main

import (
	"testing"
	"fmt"
	"strings"
	"github.com/julienschmidt/httprouter"
	"reflect"
	"com.hoover/banana/ba"
)

func countParams(path string) uint8 {
	var n uint
	for i := 0; i < len(path); i++ {
		if path[i] != ':' && path[i] != '*' {
			continue
		}
		n++
	}
	if n >= 255 {
		return 255
	}
	return uint8(n)
}

func Test_CountParams(t *testing.T) {
	fmt.Println(countParams("stu:del:212"))
}

func Test_substr(t *testing.T) {
	str := "张三李四/王二麻子"
	fmt.Println(string(str[strings.LastIndex(str, "/", )+1:]))
	fmt.Println(string(str[:strings.IndexAny(str, "/")]))
	fmt.Println(string([]rune(str)[1:3]))
	httprouter.New()
}

type Gods struct {
	a ba.Handler `method:"get" path:"/gods/new"`
}

func (g *Gods) NewGods() {
	fmt.Println("newGods")
}

func findDoc(stru interface{}) map[string]string {
	t := reflect.TypeOf(stru).Elem()
	doc := make(map[string]string)

	doc["method"] = t.Field(0).Tag.Get("method")
	doc["path"] = t.Field(0).Tag.Get("path")
	return doc

}

func Test_tag(t *testing.T) {
	var g Gods
	doc := findDoc(&g)
	fmt.Println(doc["method"], doc["path"])
}

type Fuc func(name string)

func NewObj(typ interface{}) interface{} {
	mTty := reflect.Indirect(reflect.ValueOf(typ)).Type() //获得类型
	//再造一个类型
	vc := reflect.New(mTty)
	return vc.Interface()

}

func NewFun(typ interface{}) Fuc {
	return func(name string) {
		NewObj(typ)
	}
}

func Test_Reflect_New(t *testing.T) {
	fun, ok := NewObj(&Gods{}).(Gods)
	if ok {
		fun.NewGods()
	}
}
