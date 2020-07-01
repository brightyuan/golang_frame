package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string    `json:"name"`
	Email string   `json:"email"`
	Hobby []interface{} `json:"hobby"`
}

type Response struct {
	Status int `json:"status"`
	Msg    string `json:"msg"`
	//data   interface{} `json:"data"`
}

func main() {
	u1 := Response{
		Msg: "七米",
	}
	// struct -> json string
	b, err := json.Marshal(u1)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
}