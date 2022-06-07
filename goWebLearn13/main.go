package main

import (
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/home",func(w http.ResponseWriter, r *http.Request) {
		t,_ :=template.ParseFiles("layout.html" , "home.html")
		t.ExecuteTemplate(w,"layout","Hello World")
	})
	http.HandleFunc("/contact",func(w http.ResponseWriter, r *http.Request) {
		t,_:=template.ParseFiles("layout.html")
		t.ExecuteTemplate(w,"layout","")
	})
	http.HandleFunc("/about",func(w http.ResponseWriter, r *http.Request) {
		t,_:=template.ParseFiles("layout.html","about.html")
		t.ExecuteTemplate(w,"layout","")
	})
	http.ListenAndServe("localhost:8080",nil)
}