package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jiuchengyouni/goWebLearn/tree/main/goWebLearn16/middleware"
)

func main() {
	// http.HandleFunc("/companies",func(w http.ResponseWriter, r *http.Request) {
	// 	switch r.Method{
	// 	case http.MethodPost:
	// 		dec:=json.NewDecoder(r.Body)
	// 		company:=Company{}
	// 		err:=dec.Decode(&company)
	// 		if err!=nil{
	// 			log.Println(err.Error())
	// 			w.WriteHeader(http.StatusInternalServerError)
	// 		    return
	// 		}

	// 		enc:=json.NewEncoder(w)
	// 		err=enc.Encode(company)
	// 		if err!=nil{
	// 			log.Println(err.Error())
	// 			w.WriteHeader(http.StatusInternalServerError)
	// 		    return
	// 		}
	// 	default:
	// 		w.WriteHeader(http.StatusMethodNotAllowed)
	// 	}
	// })
	http.HandleFunc("/companies",func(w http.ResponseWriter, r *http.Request) {
		c:=Company{
			ID: 123,
			Name: "珈乐",
			Country: "booboo",
		}
		enc:=json.NewEncoder(w)
		enc.Encode(c)
	})
	http.ListenAndServe("localhost:8080",new(middleware.AuthMiddleware))
	jsonStr:=`{
		"id":123,
		"name":"珈乐",
		"country":"啵啵"
	}`

	c:=Company{}
	_=json.Unmarshal([]byte(jsonStr),&c)
	fmt.Println(c)

	bytes,_:=json.Marshal(c)
	fmt.Println(string(bytes))

	bytes1,_:=json.MarshalIndent(c,""," ")
	fmt.Println(string(bytes1))
}