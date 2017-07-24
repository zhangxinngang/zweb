package handler

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"zweb/dao"
)

func GetBlogHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	blog := dao.GetBlog(r.Form.Get("id"))
	info, _ := json.Marshal(blog)
	fmt.Println(r.RequestURI, r.Method, blog)
	fmt.Fprintf(w, "Hello, %q,%v", html.EscapeString(r.URL.Path), string(info))
}

func PostBlogHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("post")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	blog := dao.Blog{}
	json.Unmarshal(data, &blog)
	dao.AddBlog(blog)
	w.Write([]byte("asdsd"))
}
