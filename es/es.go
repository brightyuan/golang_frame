package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"reflect"
	"strconv"
)

type Subject struct {
	ID     int
	Title  string
	Genres []string
}

func main() {

	var (
		indexName = "post"
		servers   = "http://localhost:32769"
		subject   Subject
	)

	const mapping = `
{
    "mappings": {
        "properties": {
            "id": {
                "type": "long"
            },
            "title": {
                "type": "text"
            },
            "genres": {
                "type": "keyword"
            }
        }
    }
}`
	ctx := context.Background()
	client := newClient(servers)
	//1. ping
	info, code, err := client.Ping(servers).Do(ctx)
	if err != nil {
		return
	}
	log.Printf("code %d and version %s\n", code, info.Version.Number)

	//2.创建索引
	createIndex(client, indexName, mapping)

	//3.写入doc数据
	subject = Subject{ID: 11, Title: "hello", Genres: []string{"jjjjj", "test"}}
	writeDoc(client, indexName, subject)

	//4.按id查询
	readDocById(client, indexName, subject)
	//5.按struct查询
	queryDoc(client, ctx, indexName, subject)
	delDoc(client, indexName, 0)

}

func newClient(servers string) *elastic.Client {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(servers))
	if err != nil {
		panic(err)
	}
	return client
}

func createIndex(client *elastic.Client, indexName string, mapping string) {
	ctx := context.Background()

	exists, err := client.IndexExists(indexName).Do(ctx)
	if err != nil {
		panic(err)
	}
	if !exists {
		_, err := client.CreateIndex(indexName).BodyString(mapping).Do(ctx)
		if err != nil {
			panic(err)
		}
		fmt.Println("创建索引成功!")
	}
	fmt.Println(exists)
}

func writeDoc(client *elastic.Client, indexName string, subject Subject) {
	ctx := context.Background()
	doc, err := client.Index().Index(indexName).Type("_doc").Id(strconv.Itoa(subject.ID)).BodyJson(subject).Refresh("wait_for").Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed with id=%v,type = %s\n", doc.Id, doc.Type)
}

func readDocById(client *elastic.Client, indexName string, subject Subject) {
	ctx := context.Background()
	//doc.Found为真表示找到这个记录了，doc.Source里面，可以手动用json.Unmarshal解析到结构体变量上
	doc, err := client.Get().
		Index(indexName).
		Id(strconv.Itoa(subject.ID)).
		Do(ctx)
	if err != nil {
		panic(err)
	}
	if doc.Found {
		fmt.Printf("Got doc  id=%v,type = %s\n", doc.Id, doc.Type)
	}
	err = json.Unmarshal(*doc.Source, &subject)
	if err != nil {
		panic(err)
	}
	fmt.Println(subject.ID, subject.Title, subject.Genres)
}

func queryDoc(client *elastic.Client, ctx context.Context, indexName string, subject Subject) {
	termQuery := elastic.NewTermQuery("title", subject.Title)
	result, err := client.Search().Index(indexName).Type("_doc").Query(termQuery).Do(ctx)
	//.Sort("id", true).From(0).Size(10).Pretty(true).Do(ctx)
	if err != nil {
		panic(err)
	}
	total := result.TotalHits()
	if total > 0 {
		for _, item := range result.Each(reflect.TypeOf(subject)) {
			if t, ok := item.(Subject); ok {
				fmt.Printf("query Found: Subject(id=%d, title=%s)\n", t.ID, t.Title)
			}
		}
	} else {
		fmt.Println("not found!")
	}

}

func delDoc(client *elastic.Client, indexName string, id int) {
	ctx := context.Background()
	res, err := client.Delete().
		Index(indexName).Type("_doc").
		Id(strconv.Itoa(id)).
		Refresh("wait_for").
		Do(ctx)
	if err != nil {
		panic(err)
	}
	if res.Result == "deleted" {
		fmt.Println("Document 1: deleted")
	}
	fmt.Println("****")
}
