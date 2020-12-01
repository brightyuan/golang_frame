package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	//Hobby []interface{} `json:"hobby"`
	Hobby Hobby `json:"hobby"`
}

type Hobby struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	//data   interface{} `json:"data"`
}

func main() {
	hobby := Hobby{
		Status: 1,
		Msg:    "hello",
	}
	u1 := User{
		Name:  "七米",
		Hobby: hobby,
	}

	// struct -> json string
	b, err := json.Marshal(u1)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Print(string(b))
	//fmt.Printf("str:%s\n", b)
}
