package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))

	http.Handle("/", fs)

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server in port 8080")

	fmt.Println(http.ListenAndServe(":8080", nil))
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "error while parsing the form "+err.Error(), http.StatusConflict)
		return
	}

	fmt.Fprintf(w, "Form processed successfully\n")

	fmt.Fprintf(w, "Name: %s\n", r.FormValue("name"))
	fmt.Fprintf(w, "Address: %s\n", r.FormValue("address"))
}

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "invalid path", http.StatusNotFound)
	}

	if r.Method != http.MethodGet {
		http.Error(w, r.Method+" not supported, only get is supported", http.StatusNotFound)
	}

	fmt.Fprintf(w, "hello")
}
