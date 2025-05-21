## How to serve static files e.g., CSS, JS, Images...

#### IMPORTANT! Go does not modify or "read" the contents of your CSS or JS — it just serves them.

Static files are files that don’t change on the server side — they are delivered to the client exactly as they are
stored.

### Using filesystem (read from disk at runtime)

```go
package main

func main() {
	// When someone requests a file under /static/, go will look inside this folder and serve that file.
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)
	})

	log.Println("Server started on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Using go:embed (files are compiled into the binary at build time)

```go
package main

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
		tmpl.Execute(w, nil)
	})

	log.Println("server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```