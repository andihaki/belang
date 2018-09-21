package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type users struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func serveJSON() {
	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var user users
		json.NewDecoder(r.Body).Decode(&user)

		fmt.Fprintf(w, "%s %s is %d years old!", user.Firstname, user.Lastname, user.Age)
	})

	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		wick := users{
			Firstname: "John",
			Lastname:  "Wick",
			Age:       40,
		}

		json.NewEncoder(w).Encode(wick)
	})

	http.ListenAndServe(":8080", nil)
}
