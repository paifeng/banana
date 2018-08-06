package ba

import (
	"net/http"
	"mime/multipart"
	"github.com/banana/bhttp"
	"strconv"
)

type reqI struct {
	Req
	req *http.Request
}

//获取请求路径
func (ri *reqI) GetPath() string {
	return ri.req.URL.Path
}

//获取指定key的请求参数值，单个值
func (ri *reqI) GetParam(key string) string {
	return ri.req.FormValue(key)
}

//获取指定key的请求参数值，int
func (ri *reqI) GetParamInt(key string) int {
	val, err := strconv.Atoi(ri.req.FormValue(key))
	if err != nil {

	}
	return val
}

//内省结构体
func (ri *reqI) BindOBJ(obj interface{}) error {
	if ri.req.Form == nil {
		ri.req.ParseMultipartForm(32 << 20) //32MB
	}
	return bhttp.MapForm(obj, ri.req.Form)
}

//获取文件
func (ri *reqI) GetFile(key string) (multipart.File, *multipart.FileHeader, error) {
	return ri.req.FormFile(key)
}
