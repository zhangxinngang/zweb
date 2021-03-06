package route

import (
	"fmt"
	"net/http"
)

type Content struct {
	Resp http.ResponseWriter
	Req  *http.Request
}

type RouteNode struct {
	Method  string
	Path    string
	Handler func(http.ResponseWriter, *http.Request)
}

var Routes = map[string][]RouteNode{}

func AddRoute(method, path string, handler func(http.ResponseWriter, *http.Request)) {
	if Routes[path] == nil {
		http.HandleFunc(path, IndexHandler)
	}
	//尽量常用的method放前面 如GET POST
	Routes[path] = append(Routes[path], RouteNode{Method: method, Path: path, Handler: handler})
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	nodes := Routes[r.URL.Path]
	for _, node := range nodes {
		if node.Method == r.Method {
			node.Handler(w, r)
			return
		}
	}
	fmt.Fprintf(w, "method not allowd,%v", r.Method)
}
