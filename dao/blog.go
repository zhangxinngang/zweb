package dao

import (
	"context"
	"fmt"

	"reflect"

	elastic "gopkg.in/olivere/elastic.v5"
)

type Blog struct {
	Id         string `json:"id"`
	Editor     string `json:"editor"`
	Blog       string `json:"blog"`
	CreateTime int64  `json:"create_time"` // second
}

func GetBlog(id string) Blog {
	//get, err := EsClient.Get().Index("blog").Type("blog").Id(id).Do(context.Background())
	//fmt.Println(get, err, get.Fields)
	searchResult, err := EsClient.Search().
		Index("blog").                         // search in index "twitter"
		Query(elastic.NewTermQuery("id", id)). // specify the query
		Sort("id", true).                      // sort by "user" field, ascending
		From(0).Size(10).                      // take documents 0-9
		Pretty(true).                          // pretty print request and response JSON
		Do(context.Background())
	fmt.Println("get err:", err)
	result := searchResult.Each(reflect.TypeOf(Blog{}))[0].(Blog)
	return result
}

func AddBlog(blog Blog) {
	resp, err := EsClient.Index().Index("blog").Type("blog").Id(blog.Id).BodyJson(blog).Do(context.Background())
	fmt.Println(resp, err)
}

func DelBlog(id string) {
	delete(GlogMap, id)
}
