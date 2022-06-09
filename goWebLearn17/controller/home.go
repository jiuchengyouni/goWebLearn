package controller

import (
	"html/template"
	"net/http"
)

func registerHomeRoutes() {
	http.HandleFunc("/home", handleHome)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	if pusher, ok := w.(http.Pusher); ok {
		pusher.Push("/css/app.css", &http.PushOptions{
			Header: http.Header{
				"Content-Type": []string{"text/css"}},
		})
	}
	t, _ := template.ParseFiles("layout.html", "home.html")
	t.ExecuteTemplate(w, "layout", "Hello World")
}
