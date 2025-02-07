package main

import (
	//"fmt"
	"database/sql"
	"html/template"
	"net/http"
)

type User struct {
	Id       int
	Name     string
	Password string
}

func home_page(w http.ResponseWriter, r *http.Request) {

	tmpl, _ := template.ParseFiles("Pages/timeTable.html")
	if tmpl != nil {
		tmpl.Execute(w, nil)
	}
}

func contact_page(w http.ResponseWriter, r *http.Request) {

}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contact_page)
	http.ListenAndServe(":8080", nil)
}

func main() {

	db, err := sql.Open("posgresql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO users(id, name, password) VALUES(1,'12','52')")

	if err != nil {
		panic(err)
	}
	defer insert.Close()

	handleRequest()

}
