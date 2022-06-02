package main

import "net/http"

type helloHandler struct{}
func(m *helloHandler)ServeHTTP(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("Hello World"))
}
type aboutHandler struct{}
func(m *aboutHandler)ServeHTTP(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("about 珈乐"))
}
func welcome(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("welcome 珈乐"))
}
func main() {
	mh:=helloHandler{}
	a:=aboutHandler{}
	//创建Web Server
	//单个handler
	//1.
	//server:=http.Server{
	//	Addr:"localhost:8080",
	//	Handler: &mh,
	//}
	//server.ListenAndServe()
	//2.http.ListenAndServe("localhost:8080",&mh)
	//多个hanlder
	server:=http.Server{
		Addr:"localhost:8080",
		Handler: nil,
	}
	//输入localhost/hello返回helloworld
	http.Handle("/hello",&mh)
	//输入localhost/about返回about
	http.Handle("/about",&a)
	//另一种写法
	//输入localhost/home返回home
	http.HandleFunc("/home",func(w http.ResponseWriter,r *http.Request){
		w.Write([]byte("gohome 珈乐"))
	})
	//还有一种写法
	http.Handle("/welcome",http.HandlerFunc(welcome))
	server.ListenAndServe()
}	
