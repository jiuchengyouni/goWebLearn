package main

import (
	"fmt"
	"net/http"
)

func main() {
	server:=http.Server{
		Addr: "localhost:8080",
	}

	http.HandleFunc("/process",func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Fprintln(w,r.Form)
	})

	server.ListenAndServe()
}