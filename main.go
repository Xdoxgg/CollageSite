package main

import (
	"html/template"
	"net/http"
	"database/sql"
    "github.com/lib/pq""database/sql"
    "github.com/lib/pq"
)


type Group struct {
    ID        int    `json:"id"`         
    GroupName string `json:"group_name"` 
}

type Students struct { 
	ID        	int	   `json:"id"` 
	GroupID  	int    `json:"group_id"`
	StudentName string `json:"student_name"`
}


type Lessons struct {
	ID	  		 int 	`json:"id"`
	Title 		 string `json:"title"`
	LessonNumber int 	`json:"lesson_number"`
	LessonDate	 int 	`jspn:"lesson_date"`

}

func connectDB() (*sql.DB, error) {
	connStr := "user= password= dbname= sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        fmt.Println("Ошибка подключения к базе данных:", err)
        return
    }
    defer db.Close()
	
	
	return sql.Open("postgres", connStr)
}

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
	
	http.ListenAndServe(":8080", nil)
}

func GetStudents(db *sql.DB) ([]Student, error) {
    rows, err := db.Query(`SELECT id, group_id, student_name FROM students`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var students []Student
    for rows.Next() {
        var student Student
        err := rows.Scan(&student.ID, &student.GroupID, &student.StudentName)
        if err != nil {
            return nil, err
        }
        students = append(students, student)
    }
    return students, nil
}

func main() {

	handleRequest()
	connectDB()


}