package http

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name       string
	Method     string
	path       string
	HandleFunc http.HandlerFunc
}

type Routes []Route

//配置路由信息
var routes = Routes{
	{
		"Index",
		"GET",
		"/api/v1",
		Index,
	},
	{
		"GetImageUsageByName",
		"GET",
		"/api/v1/ceph/monitor/rdb/image/{imageName}",
		GetImageUsageByName,
	},
}

var router *mux.Router

//初始化路由信息
func NewRouter() *mux.Router {
	if router == nil {
		router = mux.NewRouter().StrictSlash(true)
	}
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.path).
			Name(route.Name).
			Handler(route.HandleFunc)
	}
	return router
}
