package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func progess(w http.ResponseWriter,r *http.Request){
	// r.ParseMultipartForm(1024)

	// fileHeader:=r.MultipartForm.File["uploaded"][0]
	// file,err:=fileHeader.Open()
	file,_,err:=r.FormFile("uploaded")
	if err==nil{
		date,err:=ioutil.ReadAll(file)
		if err==nil{
			fmt.Fprintln(w,string(date))
		}
	}
}
func main() {
	server:=http.Server{
		Addr: "localhost:8080",
	}
	http.Handle("/process",http.HandlerFunc(progess))

	server.ListenAndServe()
}