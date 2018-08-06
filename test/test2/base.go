package main

import "github.com/banana/ba"

//type Action struct {
//	Desc    string
//	Method  string
//	Path    string
//	Handler func(req ba.Req, resp ba.Resp)
//}

type Action func(req ba.Req, resp ba.Resp)

type article struct {
	ADD Action `method:"POST" path:"/article/add" desc:"添加一篇文章"`
	DEL Action `method:"GET" path:"/article/del" desc:"删除一篇文章"`
}



