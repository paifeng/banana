package ba

import (
	"html/template"
	"mime/multipart"
)

type Req interface {
	GetPath() string
	GetParam(key string) string
	GetParamInt(key string) int
	BindOBJ(obj interface{}) error
	GetFile(key string) (multipart.File, *multipart.FileHeader, error)
}
type Resp interface {
	String(status int, body string)
	Json(status int, obj interface{})
	HTML(status int, templateAlias string, data interface{})
	Redirect(location string)
}

//回调Handle
type Handler func(req Req, resp Resp)

type Controller interface{}

//主接口
type Engine interface {
	Get(path string, controller interface{})
	Post(path string, handler Handler)
	Controller(c Controller)
	RouterToString() string
	Run()
	AddTemplate(alias string, filename string)
	Template() (*template.Template)
}
