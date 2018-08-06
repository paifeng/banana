package ba

import (
	"radix-tree"
)

type Router struct {
	Trees *iradix.Tree //基数树，用于存储路由
}

//路由处理方法，注册回调
func (ri *Router) Handle(method, path string, controller interface{}) bool {
	if path[0] != '/' { //路径必须以'/'开头
		panic("path must begin with '/' in path ['" + path + "']")
	}
	var ok bool
	ri.Trees, _, ok = ri.Trees.Insert([]byte(path), controller)
	return ok
}

//转换成map
func (ri *Router) ToMap() map[string]interface{} {
	return ri.Trees.ToMap()
}
