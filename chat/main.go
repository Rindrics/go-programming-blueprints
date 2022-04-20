package main

import (
	"log"
	"net/http"
	"sync"
	"text/template"
)

type templateHandler struct {
	filename string
	once     sync.Once
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<!DOCTYPE html>
				<html>
					<head>
						<title>Chat</title>
					</head>
					<body>
						Let's chat!
					</body>
				</html>

		`))
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
