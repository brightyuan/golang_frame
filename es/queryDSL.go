package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
)

func main() {
	query := elastic.NewTermQuery("Title", "hello")
	src, err := query.Source()
	if err != nil {
		panic(err)
	}
	PrintQuery(src)
	servers := "http://localhost:32769"

	client := newClient(servers)
	indexName := "post"
	ctx := context.Background()
	result, err := client.Search().Index(indexName).Query(query).Do(ctx)
	fmt.Println(result.TotalHits())
}

func newClient(servers string) *elastic.Client {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(servers))
	if err != nil {
		panic(err)
	}
	return client
}

func PrintQuery(src interface{}) {
	data, err := json.MarshalIndent(src, "", "   ")
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(data))
}
