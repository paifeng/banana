package ba

import (
	"net/http"
	"encoding/json"
	"html/template"
)

type respI struct {
	Resp
	writer http.ResponseWriter
	templates  *template.Template //模板对象
}

//输出字符串
func (ri *respI) String(status int, body string) {
	ri.writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	ri.writer.WriteHeader(status)
	ri.writer.Write([]byte(body))
}

//输出JSON
func (ri *respI) Json(status int, obj interface{}) {
	ri.writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	ri.writer.WriteHeader(status)
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	ri.writer.Write(jsonBytes)
}

//输出HTML模板
func (ri *respI) HTML(status int, templateAlias string, data interface{}) {
	ri.writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	ri.writer.WriteHeader(status)
	err := ri.templates.ExecuteTemplate(ri.writer, templateAlias, data)
	if err != nil {
		panic(err)
	}
}

//重定向
func (ri *respI) Redirect(location string) {
	ri.writer.Header().Set("Location", location)
	ri.writer.WriteHeader(http.StatusMovedPermanently)
}
