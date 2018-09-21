sumber:
https://gowebexamples.com
#typo
https://gowebexamples.com/templates/
pares => parse



# http
http.ResponseWriter = text/html yang mau ditampilin ke user (web browser)
http.Request = semua informasi di HTTP request termasuk URL

# gorilla-mux
salah satu fiturnya bisa ekstrak URL segmen
/books/go-programming-blueprint/page/10

```
r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    title := vars["title"]
    page := vars["page"]

    fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
})
```

# html/template
[godoc]https://golang.org/pkg/text/template/#hdr-Actions

```
tmpl, err := template.ParseFiles("layout.html")
// or
tmpl := template.Must(template.ParseFiles("layout.html"))
```

Control Structure	Definition
```
{{/* a comment */}}	Defines a comment
{{.}}	Renders the root element
{{.Title}}	Renders the “Title”-field in a nested element
{{if .Done}} {{else}} {{end}}	Defines an if-Statement
{{range .Todos}} {{.}} {{end}}	Loops over all “Todos” and renders each using {{.}}
{{block "content" .}} {{end}}	Defines a block with the name “content”
```

## Execute
nerima io.Writer untuk ditulis ke template