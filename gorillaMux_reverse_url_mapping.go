package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func ArticleHandler2(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Category is: %v\n", vars["category"])
	fmt.Fprintf(w, "ID is: %v\n", vars["id"])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler2).
					Name("articleRoute")

	// srv := &http.Server {
	// 	Handler: r,
	// 	Addr: "127.0.0.1:8080",
	// 	WriteTimeout: 15*time.Second,
	// 	ReadTimeout: 15 * time.Second,
	// }

	// log.Fatal(srv.ListenAndServe())

	url, err := r.Get("articleRoute").URL("category", "category_name", "id", "2")
	fmt.Println(url)
	fmt.Println(err)
}