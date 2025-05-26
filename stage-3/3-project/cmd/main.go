package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	tmpl    *template.Template
	addr    = flag.String("addr", ":8080", "HTTP listen address")
	baseURL = flag.String("baseURL", "http://localhost:8080", "Base URL")
)

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request received: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Println("Request completed successfully")
	})
}

func loadTemplates() *template.Template {
	tmpl := template.New("")
	var files []string

	// Gather .html files
	err := filepath.Walk("assets/templates", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".html") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		log.Fatal("Error walking template files:", err)
	}

	tmpl, err = tmpl.ParseFiles(files...)
	if err != nil {
		log.Fatal("Error parsing templates:", err)
	}

	return tmpl
}

func loadPages(name string, w http.ResponseWriter, _ *http.Request) {
	tmpl = loadTemplates()

	data := struct {
		Title   []string
		Message []string
	}{
		Title:   []string{"Home page header", "About page header", "Blog page header"},
		Message: []string{"MSG HOME", "MSG ABOUT", "MSG BLOG"},
	}

	if err := tmpl.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, "error executing template", http.StatusInternalServerError)
		log.Println("template execute error:", err)
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	loadPages("home.html", w, r)
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	loadPages("about.html", w, r)
}

func blogPage(w http.ResponseWriter, r *http.Request) {
	loadPages("blog.html", w, r)
	if r.Method == http.MethodPost {
		fmt.Fprint(w, `<img src="./assets/images/gopher.png" alt="sample image" width="50" height="50">`)
	}
}

func routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/assets/static/css/", http.StripPrefix("/assets/static/css/",
		http.FileServer(http.Dir("./assets/static/css"))))

	mux.Handle("/assets/images/", http.StripPrefix("/assets/images/",
		http.FileServer(http.Dir("./assets/images/"))))

	mux.Handle("/", logMiddleware(http.HandlerFunc(homePage)))
	mux.Handle("/about", logMiddleware(http.HandlerFunc(aboutPage)))
	mux.Handle("/blog", logMiddleware(http.HandlerFunc(blogPage)))

	server := &http.Server{
		Addr:         *addr,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  20 * time.Second,
	}

	fmt.Println("server started at " + *baseURL)
	log.Fatal(server.ListenAndServe())

	return mux
}

func main() {
	routes()
}
