// belajaran aja
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

// tiap item dari todo
type todo struct {
	Title string
	Done  bool
}

// array todo
type todoPageData struct {
	PageTitle string
	Todos     []todo
	Success   bool
}

// cookies
var (
	// key must 16, 24 / 32 bytes long(AES-128, 192, 256)
	key   = []byte("rahasia-boss")
	store = sessions.NewCookieStore(key)
)

func main() {
	serveHTML()
}

func serveMux() {
	// template TODO list

	r := mux.NewRouter()
	// before = http.HandleFunc, after = r.HandleFunc
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]
		fmt.Fprintf(w, "Hellow, you've requested books: %s on page %s\n", title, page)
	})

	// buat nyimpen file css, js
	fs := http.FileServer(http.Dir("static/"))
	r.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", r)
}

func serveHTML() {
	// test
	http.HandleFunc("/secret", logging(secret))
	http.HandleFunc("/login", logging(login))
	http.HandleFunc("/logout", logging(logout))

	http.HandleFunc("/", logging(fTodo))
	serverStaticFiles()
	http.ListenAndServe(":80", nil)
	fmt.Println("server running at port 80")
}

func fTodo(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("layout.html"))
	data := todoPageData{
		Success:   false,
		PageTitle: "PEER",
		Todos: []todo{
			{Title: "Belajar Golang", Done: false},
			{Title: "Create Report", Done: true},
			{Title: "Fix Bug", Done: true},
		},
	}
	if r.Method != http.MethodPost {
		tmpl.Execute(w, data)
		return
	}
	data = todoPageData{
		Success: true,
		Todos:   append(data.Todos, todo{Title: r.FormValue("title"), Done: false}),
	}
	tmpl.Execute(w, data)
}

func serverStaticFiles() {
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}

// middleware
func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

// cookie
func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "go-cookie")

	// udah login?
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden: ", http.StatusForbidden)
		return
	}

	// print secret message
	fmt.Fprintln(w, "The Cake is lie")
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "go-cookie")

	// logika autentikasi
	// ...

	// autentikasi ok
	session.Values["authenticated"] = true
	session.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "go-cookie")

	// logika autentikasi
	// ...

	// buang autentikasi
	session.Values["authenticated"] = false
	session.Save(r, w)
}
