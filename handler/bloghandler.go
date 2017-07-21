package handler

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"zweb/dao"
)

func GetBlogHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	blog := dao.GetBlog(r.Form.Get("id"))
	info, _ := json.Marshal(blog)
	fmt.Println(r.RequestURI, r.Method, blog)
	fmt.Fprintf(w, "Hello, %q,%v", html.EscapeString(r.URL.Path), string(info))
	fmt.Println("get")
}

func PostBlogHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("post")
}
