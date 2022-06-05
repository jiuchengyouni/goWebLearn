package main

import (
	"log"
	"net/http"
	"text/template"
)

//加载模板
func main() {
	templates := loadTemplates()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//切片，去除“/”
		fileName := r.URL.Path[1:]
		t := templates.Lookup(fileName)
		if t != nil {
			err := t.Execute(w, nil)
			if err != nil {
				log.Fatalln(err.Error())
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})
	//处理静态加载文件
	http.Handle("/css/", http.FileServer(http.Dir("wwwroot")))
	http.Handle("/img/", http.FileServer(http.Dir("wwwroot")))
	http.ListenAndServe("localhost:8080", nil)
}
func loadTemplates() *template.Template {
	result := template.New("templates")
	template.Must(result.ParseGlob("templates/*.html"))
	return result
}
