package main

import (
	//"fmt"
	//"database/sql"
	 "html/template"
	"net/http"
)

// type User struct {
// 	Id       int
// 	Name     string
// 	Password string
// }

func home_page(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("Pages/index.html")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, nil)
	
}



func handleRequest() {
	http.HandleFunc("/", home_page)
	
	http.ListenAndServe(":8080", nil)
}

func main() {

	handleRequest()

}
