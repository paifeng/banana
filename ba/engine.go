package ba

import (
	"net/http"
	"html/template"
	"radix-tree"
	"strings"
	"reflect"
	"fmt"
	"encoding/json"
)

//常量
const (
	version       = "1.0.0" // VERSION 当前的程序版本
	serverName    = "banana"
	templateLeft  = "{{"
	templateRight = "}}"
)

//主接口的实现
type engineI struct {
	Engine
	ServerName    string             //服务名称
	Router        *Router            //路由
	templates     *template.Template //模板对象
	templateLeft  string
	templateRight string
}

//初始化函数
func Init() Engine {
	ei := &engineI{
		ServerName:    serverName,
		templateLeft:  templateLeft,
		templateRight: templateRight,
	}
	router := &Router{
		Trees: iradix.New(), //初始化树
	}
	ei.Router = router
	ei.templates = template.New("tpl").Delims(ei.templateLeft, ei.templateRight)
	return ei
}

//运行函数
func (ei engineI) Run() {
	http.ListenAndServe(":9314", ei)
}

// ServeHTTP 实现http的方法
func (ei engineI) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	_, fun, ok := ei.Router.Trees.Root().LongestPrefix([]byte(request.URL.Path)) //找到最大匹配的URL
	if ok {
		mReq := &reqI{req: request}
		mResp := &respI{writer: w, templates: ei.templates}
		mf := fun.(func(Req, Resp))
		mf(mReq, mResp)
	} else {
		//路径不存在
		w.Write([]byte("[404]路径不存在"))
	}
}

//添加模板
func (ei engineI) AddTemplate(alias string, filename string) {
	var err error
	_, err = ei.templates.ParseFiles(filename)
	if err != nil {
		panic(err)
	}
}

//获取模板
func (ei engineI) Template() (*template.Template) {
	return ei.templates
}

//Get请求
func (ei engineI) Get(path string, controller interface{}) {
	ei.Router.Handle("GET", path, controller)
}

//Post请求
func (ei engineI) Post(path string, handler Handler) {
	ei.Router.Handle("POST", path, handler)
}

//添加控制器
func (ei engineI) Controller(c Controller) {
	if c == nil {
		panic("[error] Controller can not be nil") //控制器不能是nil
	}
	tof := reflect.TypeOf(c)  //获取结构体字段类型
	vof := reflect.ValueOf(c) //字段值
	tofEle := tof.Elem()
	numOfFields := tofEle.NumField()
	if numOfFields == 0 {
		panic("[error] Controller format error") //控制器的结构体不能没有字段
	}
	for i := 0; i < numOfFields; i++ {
		fieldITyp := tofEle.Field(i)
		fieldIVal := vof.Elem().Field(i)
		if !fieldIVal.CanSet() {
			panic("[error] the field " + fieldITyp.Name + " must can be set")
		}
		name := strings.ToUpper(fieldITyp.Name)
		method := fieldITyp.Tag.Get("method")
		if method == "" {
			panic("[error] the Filed of Controller must have tag \"method\"") //必须要有method标签
		}
		path := fieldITyp.Tag.Get("router")
		if path == "" {
			panic("[error] the Filed of Controller must have tag \"router\"") //必须要有router标签
		}
		_, ok := tof.MethodByName(name)
		if !ok {
			panic("[error] Controller must implement function \"" + name + "(ba.Req, ba.Resp)\"")
		}
		fmt.Println(name)
		fun := vof.MethodByName(name)
		if fun.CanInterface() {
			funI := fun.Interface()
			handler, ok := funI.(func(Req, Resp))
			fmt.Println(reflect.TypeOf(funI))
			if !ok {
				panic("handler type must be \"func(ba.Req, ba.Resp)\"")
			}
			//注册路由
			fmt.Println(method, path, reflect.TypeOf(handler))
			ei.Router.Handle(method, path, handler)
			fmt.Println(ok)
		}
	}
}

//输出string
func (ei engineI) RouterToString() string {
	b, err := json.Marshal(ei.Router.ToMap())
	if err != nil {
		panic(err)
	}
	return string(b)
}
