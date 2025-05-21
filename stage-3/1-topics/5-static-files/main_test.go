package main

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	//r := httptest.NewRequest("GET", "/", nil) // testing root
	r := httptest.NewRequest("GET", "/static/css/style.css", nil) // testing a specific file
	w := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFS(content, "templates/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	})

	handler.ServeHTTP(w, r)

	result := w.Result()
	if result.StatusCode != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			result.StatusCode, http.StatusOK)
	}

	body := w.Body.String()
	if !strings.Contains(body, "<title>") {
		t.Errorf("HTML content not found in response body: got %v", body)
	}
}
