package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"zweb/dao"
)

func GetBlogHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	blog := dao.GetBlog(r.Form.Get("id"))
	info, _ := json.Marshal(blog)
	fmt.Println(r.RequestURI, r.Method, blog)
	fmt.Fprintf(w, string(info))
}

func PostBlogHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("post")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	blog := dao.Blog{}
	err = json.Unmarshal(data, &blog)
	fmt.Println(err, "err", string(data))
	dao.AddBlog(blog)
	w.Write([]byte("ok"))
}
