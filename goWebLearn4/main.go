package main

import (
	"fmt"
	"net/http"
)

func main() {
	// server:=http.Server{
	// 	Addr: "localhost:8080",
	// }

	// http.HandleFunc("/url",func(w http.ResponseWriter, r *http.Request) {
	//  	fmt.Fprintln(w,r.URL.Fragment)
	// })
	// server.ListenAndServe()
	//Request Header例子
	// server:=http.Server{
	// 	Addr: "localhost:8080",
	// }
	// http.HandleFunc("/header",func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w,r.Header)
	// 	fmt.Fprintln(w,r.Header["Accept-Encoding"])
	// 	fmt.Fprintln(w,r.Header.Get("Accept-Encoding"))
	// })
	server:=http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/post",func(w http.ResponseWriter, r *http.Request) {
		length:=r.ContentLength
		body:=make([]byte,length)
		r.Body.Read(body)
		fmt.Fprintln(w,string(body))
	})
	server.ListenAndServe()
}