package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/validation"
)

type OojStudent2 struct {
	Id     int      `form:"id"`
	Name   string   `form:"name"`
	Age    int      `form:"age"`
	Gender string   `form:"gender"`
	Hobby  []string `form:"hobby"`
}

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "index.tpl"
}

//数据验证
func Valid()  {
	obj := OojStudent2{}

	valid := validation.Validation{}
	valid.Required(&obj,"OojStudent2")
}

func main() {
	beego.SetLevel(beego.LevelEmergency)
	beego.BConfig.WebConfig.TemplateLeft="<<<"
	beego.Get("/info", func(c *context.Context) {
		var id int
		var name string
		var age int
		var gender string
		var hobby []string
		c.Input.Bind(&id, "id")
		c.Input.Bind(&name, "name")
		c.Input.Bind(&age, "age")
		c.Input.Bind(&gender, "gender")
		c.Input.Bind(&hobby, "hobby")

		c.Output.JSON(&OojStudent2{id, name, age, gender, hobby}, false, false)
	})
	beego.Run(":7890")
}
