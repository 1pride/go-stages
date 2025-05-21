package main

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

//go:embed static/* templates/*
var content embed.FS

type Greetings struct {
	Msg  string
	Name string
}

func handler(w http.ResponseWriter, r *http.Request) {
	data := Greetings{
		Name: "John",
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data.Msg = r.FormValue("message")
		data.Name = r.FormValue("name")
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	if err := tmpl.Execute(w, data); err != nil {
		log.Fatal("error at executing template", err)
	}
}

func main() {
	staticFiles, err := fs.Sub(content, "static")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFiles))))

	http.HandleFunc("/", handler)
	fmt.Println("server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
