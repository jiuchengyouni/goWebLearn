package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

//条件输出
func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	//生成随机数
	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 5)
}

//有东西遍历
func process2(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl2.html")
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu"}
	t.Execute(w, daysOfWeek)
}

//没东西遍历
func process3(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl2.html")
	daysOfWeek := []string{}
	t.Execute(w, daysOfWeek)
}

//对html原先文本添加
func process4(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl3.html")
	t.Execute(w, "hello")
}
func process5(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("t1.html", "t2.html")
	t.Execute(w, "hello world")
}
func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/process4",process4)
	http.HandleFunc("/process3",process3)
	http.HandleFunc("/process2", process2)
	http.HandleFunc("/process", process)
	http.HandleFunc("/process5",process5)
	server.ListenAndServe()
}
