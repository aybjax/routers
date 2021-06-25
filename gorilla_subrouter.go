package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func ArticleHandler4(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Category is: %v\n", vars["category"])
	fmt.Fprintf(w, "ID is: %v\n", vars["id"])
}

func main() {
	r := mux.NewRouter()
	// r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler4)
	// r.Path("/articles/{category}/{id:[0-9]+}").HandlerFunc(ArticleHandler4)

	s := r.PathPrefix("/articles").Subrouter()
	s.HandleFunc("/{category}/{id:[0-9]+}", ArticleHandler4)

	srv := &http.Server {
		Handler: r,
		Addr: "127.0.0.1:8080",
		WriteTimeout: 15*time.Second,
		ReadTimeout: 15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}