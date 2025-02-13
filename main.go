package main

import (
	"html/template"
	"net/http"
)

// func timeTable(w http.ResponseWriter, r *http.Request) {
// 	tmpl, err := template.ParseFiles("Pages/timeTable.html")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	tmpl.Execute(w, nil)
// }

func index(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("Pages/index.html")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, nil)
	tmpl.ExecuteTemplate(w, "index", nil)
}

func handleRequest() {
	http.Handle("/Styles/", http.StripPrefix("/Styles/", http.FileServer(http.Dir("./Styles/"))))
	http.Handle("/Pages/", http.StripPrefix("/Pages/", http.FileServer(http.Dir("./Pages/"))))
	http.HandleFunc("/", index)
	//http.HandleFunc("/timeTable", timeTable)
	http.ListenAndServe(":8080", nil)
}

func main() {

	handleRequest()

}