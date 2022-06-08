package controller

import (
	"html/template"
	"net/http"
)
func registerAboutRoutes(){
	http.HandleFunc("/about",handleAbout)
}
func handleAbout(w http.ResponseWriter, r *http.Request) {
	t,_:=template.ParseFiles("layout.html","about.html")
		t.ExecuteTemplate(w,"layout","")
}