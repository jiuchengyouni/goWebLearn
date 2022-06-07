package main

import (
	"net/http"
	"text/template"
	"time"
)

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}
func process(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("t1.html").Funcs(funcMap)
	t.ParseFiles("t1.html")
	t.Execute(w, time.Now())
}
func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
