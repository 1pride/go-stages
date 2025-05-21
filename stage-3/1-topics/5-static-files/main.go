package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
)

//go:embed static/* templates/*
var content embed.FS

func main() {
	staticFiles, err := fs.Sub(content, "static")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFiles))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFS(content, "templates/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			fmt.Println("error at executing template:", err)
			return
		}
	})

	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, err.Error(), http.StatusMethodNotAllowed)
			return
		}

		type Feedback struct {
			Name    string `json:"name"`
			Message string `json:"message"`
		}

		var fb Feedback
		fbJSONFile := "data/feedback.json"
		var fbList []Feedback

		err := json.NewDecoder(r.Body).Decode(&fb)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Printf("Received feedback: %+v\n", fb)

		if data, err := os.ReadFile(fbJSONFile); err == nil && len(data) > 0 {
			if err := json.Unmarshal(data, &fbList); err != nil {
				fmt.Println("error at unmarshalling:", err)
			}
		}

		fbList = append(fbList, fb)

		newData, _ := json.MarshalIndent(fbList, "", " ")
		if err := os.WriteFile(fbJSONFile, newData, 0644); err != nil {
			fmt.Println("error at writing file:", err)
		}

		w.WriteHeader(http.StatusOK)
	})

	log.Println("server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
