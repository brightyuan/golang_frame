package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:9000/hello")
	if err != nil {
		panic("err connect")
	}
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("%s",body)
}
