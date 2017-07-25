package dao

import (
	"context"
	"fmt"

	elastic "gopkg.in/olivere/elastic.v5"
)

var GlogMap map[string]Blog
var EsClient *elastic.Client

func init() {
	esClient, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		fmt.Println(err)
	}
	EsClient = esClient
	BlogMapping := `
{
	"settings":{
		"number_of_shards":1,
		"number_of_replicas":0
	},
	"mappings":{
		"_default_": {
			"_all": {
				"enabled": true
			}
		},
		"blog":{
			"properties":{
				"id":{
					"type":"keyword"
				},
				"blog":{
					"type":"text",
					"store": true,
					"fielddata": true
				},
				"create_time":{
					"type":"long"
				},
				"editor":{
					"type":"keyword"
				}
			}
		}
	}
}
`
	exists, _ := esClient.IndexExists("blog").Do(context.Background())
	if !exists {
		result, err := esClient.CreateIndex("blog").Body(BlogMapping).Do(context.Background())
		fmt.Println(result, err)
	}
}
