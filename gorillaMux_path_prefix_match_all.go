package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/static").
				Handler(http.StripPrefix("/static", http.FileServer(http.Dir("staticter"))))

	srv := &http.Server {
		Handler: r,
		Addr: "127.0.0.1:8080",
		WriteTimeout: 15*time.Second,
		ReadTimeout: 15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}