package main

import (
	"net/http"
	"github.com/jiuchengyouni/goWebLearn/tree/main/goWebLearn17/controller"
)
//访问https，自带证书
func main() {
	controller.RegisterRoutes()
	http.ListenAndServeTLS("localhost:8080","cert.pem","key.pem",nil)
}