package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter,r *http.Request){
	t,_:=template.ParseFiles("tmpl.html")
	t.Execute(w,"Hello World!")
}
func main() {
	server:=http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/process",process)
	server.ListenAndServe()
}