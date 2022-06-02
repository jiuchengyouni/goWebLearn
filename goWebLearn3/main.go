package main

import "net/http"

func main() {
	//五个内置的handler
	http.ListenAndServe(":8080",http.FileServer(http.Dir("wwwroot")))
	//http.StripPrefix
	//http.TimeoutHandler
	//http.FileServer
}