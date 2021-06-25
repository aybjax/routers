package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Server(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	io.WriteString(w, "hello world!\n")
}

func main() {
	http.HandleFunc("/hello", Server)
	log.Fatal(http.ListenAndServe(":8080", nil))
}