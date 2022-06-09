package controller

import (
	"html/template"
	"net/http"
)
func registerContactRoutes(){
	http.HandleFunc("/contact",handleContact)
}

func handleContact(w http.ResponseWriter, r *http.Request) {
	t,_:=template.ParseFiles("layout.html")
	t.ExecuteTemplate(w,"layout","")
}