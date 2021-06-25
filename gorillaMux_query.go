package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Handler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	r := mux.NewRouter()
	r.Host("aa.bb.cc")
	r.Schemes("http", "https")
	r.Methods("GET", "post", "Put")
	r.HandleFunc("/{id}", Handler)
}