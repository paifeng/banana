package main

import (
	"net/http"
	"github.com/banana/ba"
	"runtime"
	"fmt"
	"time"
)

type teacher struct {
	Say  int `method:"get" router:"/say"`
	Info int `method:"get" router:"/info"`
	Del  int `method:"get" router:"/del"`
	Add  int `method:"get" router:"/add"`
}

type OojTeacher struct {
	Id     int      `form:"id"`
	Name   string   `form:"name"`
	Age    int      `form:"age"`
	Gender string   `form:"gender"`
	Hobby  []string `form:"hobby"`
}

func (s *teacher) SAY(req ba.Req, resp ba.Resp) {
	t := &OojTeacher{
		Id:     req.GetParamInt("id"),
		Name:   req.GetParam("name"),
		Age:    req.GetParamInt("age"),
		Gender: req.GetParam("gender"),
	}
	resp.Json(http.StatusOK, t)
}

func (s *teacher) INFO(req ba.Req, resp ba.Resp) {
	t := &OojTeacher{}
	err := req.BindOBJ(t)
	if err != nil {
		resp.String(http.StatusOK, err.Error())
		return
	}
	resp.Json(http.StatusOK, t)
}

func (s *teacher) DEL(_ ba.Req, resp ba.Resp) {
	resp.String(http.StatusOK, "DEL")
}

func (s *teacher) ADD(_ ba.Req, _ ba.Resp) {
	var i int
	for i = 0; i < 10; i++ {
		i++
		time.Sleep(1000)
	}
	fmt.Println(i)
}


func index(_ ba.Req, resp ba.Resp) {
	resp.String(http.StatusOK, "index")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	b := ba.Init()

		b.Controller(&teacher{})
		//路由
		b.Get("/index", index)
		fmt.Println(b.RouterToString())
		//模板
		//b.AddTemplate("index", "C:\\Users\\love5\\Desktop\\index.html")
		//b.AddTemplate("details", "C:\\Users\\love5\\Desktop\\details.html")
		b.Run()
	}
