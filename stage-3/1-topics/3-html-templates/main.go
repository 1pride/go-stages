package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type PageData struct {
	Title  string
	People []Person
}

type Person struct {
	Name    string
	Age     int
	Above21 bool
}

func handler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "WELCOME PAGE",
		People: []Person{
			{Name: "John", Age: 20, Above21: false},
			{Name: "Jane", Age: 25, Above21: true},
			{Name: "Jake", Age: 30, Above21: true},
		},
	}

	//funcMap := template.FuncMap{
	//	"upper": strings.ToUpper,
	//}
	//
	//tmpl := template.Must(template.New("index.html").Funcs(funcMap).ParseFiles("templates/layout.html",
	//	"templates/index.html"))

	tmpl := template.Must(template.ParseFiles(
		"templates/pages/layout.html",
		"templates/partials/index.html",
	))

	if err := tmpl.Execute(w, data); err != nil {
		log.Fatal("error executing ", err)
	}
}
func main() {
	// Test in the console
	tmpl := template.Must(template.New("test").Parse("Hello, {{ .Name }}!"))
	data := map[string]string{"Name": "John"}
	tmpl.Execute(os.Stdout, data)

	http.HandleFunc("/", handler)
	fmt.Println("\nserver started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
