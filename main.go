package main

import (
	"net/http"
	"zweb/handler"
	"zweb/route"
)

func main() {
	route.AddRoute("GET", "/main", handler.GetBlogHandler)
	route.AddRoute("POST", "/main", handler.PostBlogHandler)
	http.ListenAndServe("127.0.0.1:8000", nil)
}
