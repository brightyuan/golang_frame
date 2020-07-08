package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "您看到我了")
}
func sayHello2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "您看到word")
}

func main() {
	http.HandleFunc("/hello", sayHello)
	http.HandleFunc("/world", sayHello2)
	log.Println("启动了")
	err := http.ListenAndServe("localhost:9000", nil)
	if err != nil {
		log.Fatal("List 9000")
	}
}
