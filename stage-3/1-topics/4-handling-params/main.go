package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Accessing query params
//func handler(w http.ResponseWriter, r *http.Request) {
//	query := r.URL.Query() // parsing query params
//
//	searchTerm := query.Get("query") // get first value for that key
//	page := query.Get("page")
//
//	fmt.Println(query["query"], query["page"]) // access all values
//
//	fmt.Fprintf(w, "search term: %s, page: %s", searchTerm, page)
//}

func handler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm() // parses both URL query parameters and form data into r.Form.
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	colors := r.Form["colors"] // get all values for that key

	username := r.FormValue("username") // get first value for that key
	password := r.FormValue("password") // r.FormValue(key) is a shortcut for getting the first value from r.Form.

	tmpl, err := template.ParseGlob("templates/**/*.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = tmpl.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "username: %s, password: %s, color: %v", username, password, colors)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
