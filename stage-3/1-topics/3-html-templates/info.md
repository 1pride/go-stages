## Syntax usages

| Feature          | Syntax                            |
|------------------|-----------------------------------|
| Variable         | `{{ .Field }}`                    |
| If Statement     | `{{ if .Condition }}...{{ end }}` |
| Looping          | `{{ range .Items }}...{{ end }}`  |
| Custom Funcs     | `.Funcs(funcMap)`                 |
| Template Nesting | `{{ template "name" . }}`         |
| Template Block   | `{{ define "name" }}`             |

### Custom Funcs Example

```go
funcMap := template.FuncMap{
"upper": strings.ToUpper,
}

tmpl := template.Must(template.New("index.html").Funcs(funcMap).ParseFiles("templates/index.html"))
```

and in html you can use like this:

```html
<p>Hello, {{ upper .Name }}!</p>
```

### Nesting Template Organization

```gotemplate
// Call/Insert the block defined as "header"
{{ template "header" . }}

// define a content code block called "header" being able to use in other parts
{{ define "header" }}
code block
{{ end }}
```

```
 templates/
├── layout.html
├── partials/
│   ├── header.html
│   ├── footer.html
│   └── navbar.html
└── pages/
├── index.html
└── about.html
```

in layout.html you can use the templates like:

```gotemplate
{{ template "header" . }}
{{ template "navbar" . }}
{{ template "content" . }}
{{ template "footer" . }}
```

and each html inside partials need to define templates

```html
{{ define "header" }}
<header><h1>Random Header</h1></header>
{{ end }}
```

then in go code you parse all files

```go
tmpl := template.Must(template.ParseFiles(
"templates/layout.html",
"templates/partials/header.html",
"templates/partials/footer.html",
"templates/partials/navbar.html",
"templates/pages/index.html",
))
```

or even a Glob to parse recursively html files e.g.

```go
template.ParseGlob("templates/**/*.html")
```