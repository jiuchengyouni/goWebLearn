package main

import (
	"net/http"
	"github.com/jiuchengyouni/goWebLearn/tree/main/goWebLearn15/controller"
)

func main() {
	controller.RegisterRoutes()
	http.ListenAndServe("localhost:8080",nil)
}