package main

import (
	"fmt"
	"github.com/banana/ba"
	"runtime"
	"net/http"
)

type objTeacher struct {
	Id     int      `form:"id"`
	Name   string   `form:"name"`
	Age    int      `form:"age"`
	Gender string   `form:"gender"`
	Hobby  []string `form:"hobby"`
}



var ArticleI = &article{
	ADD: func(req ba.Req, resp ba.Resp) {
		t := &objTeacher{}
		err := req.BindOBJ(t)
		if err != nil {
			resp.String(http.StatusOK, err.Error())
			return
		}
		resp.Json(http.StatusOK, t)
	},
	DEL: func(req ba.Req, resp ba.Resp) {
		t := &objTeacher{
			Id:     req.GetParamInt("id"),
			Name:   req.GetParam("name"),
			Age:    req.GetParamInt("age"),
			Gender: req.GetParam("gender"),
		}
		resp.Json(http.StatusOK, t)
	},
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	b := ba.Init()
	b.Controller(ArticleI)
	fmt.Println(b.RouterToString())
	b.Run()
}
