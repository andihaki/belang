package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	serve()
}

func serve() {
	r := mux.NewRouter()
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]
		fmt.Fprintf(w, "Hellow, you've requested books: %s on page %s\n", title, page)
	})

	// buat nyimpen file css, js
	fs := http.FileServer(http.Dir("static/"))
	r.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", r)
}
