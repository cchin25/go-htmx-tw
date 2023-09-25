package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting server...")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, nil)
	}

	sayhi := func(w http.ResponseWriter, r *http.Request) {
		name := r.PostFormValue("name")
		hellur := fmt.Sprintf("Hello %s", name)
		tmpl, _ := template.New("t").Parse(hellur)
		tmpl.Execute(w, nil)
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/sayhi", sayhi)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
