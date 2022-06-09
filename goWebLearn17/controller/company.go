package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
)

func registerCompanyRoutes() {
	http.HandleFunc("/companies", handleCompanies)
	http.HandleFunc("/companies/", handleCompany)
}
func handleCompanies(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("layout.html", "companies.html")
	t.ExecuteTemplate(w, "layout", nil)
}
func handleCompany(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("layout.html", "company.html")
	//正则表达式
	pattern, _ := regexp.Compile(`/companies/(\d+)`)
	matches := pattern.FindStringSubmatch(r.URL.Path)
	//matches>0说明找到符合
	if len(matches) > 0 {
		fmt.Println(matches[0])
		companyID, _ := strconv.Atoi(matches[1])
		t.ExecuteTemplate(w, "layout", companyID)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
