package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq" // анонимный импорт
	"html/template"
	"net/http"
)

type News struct {
	ID       int    `json:"id"`
	Data     string `json:"data"`
	Img      string `json:"img"`
	PostDate string `json:"post_date"`
	Title    string `json:"title"`
}

type Group struct {
	ID        int    `json:"id"`
	GroupName string `json:"name"`
}

type Student struct {
	ID          int    `json:"id"`
	StudentDate string `json:"student_date"`
	GroupID     int    `json:"group_id"`
	StudentName string `json:"name"`
}

type Teacher struct {
	ID   int    `json:"id"`
	name string `json:"name"`
}

type Lesson struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	LessonNumber int    `json:"lesson_number"`
	LessonDay    int    `json:"lesson_day"`
	Place        int    `json:"place"`
	GroupID      int    `json:"group_id"`
}

type mark struct {
	ID       int    `json:"id"`
	Value    string `json:"value"`
	Disciple string `json:"disciple"`
	Date     string `json:"date"`
}

func connectDB() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Ошибка подключения к базе данных:", err)
		return nil, err
	}
	// Проверяем соединение
	if err = db.Ping(); err != nil {
		fmt.Println("Не удалось подключиться к базе данных:", err)
		return nil, err
	}
	return db, nil
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("Pages/index.html")
	if err != nil {
		http.Error(w, "Ошибка загрузки шаблона", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func getGroupsHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		http.Error(w, "Ошибка подключения к базе данных", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	groups, err := getGroups(db) // Получаем данные групп из базы
	if err != nil {
		http.Error(w, "Ошибка получения данных групп", http.StatusInternalServerError)
		return
	}

	// Преобразуем данные в JSON и отправляем клиенту
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groups)
}

func getGroups(db *sql.DB) ([]Group, error) {

	rows, err := db.Query(`
	SELECT id, group_name 
	FROM groups
	UNION
	SELECT id, teacher_name
	FROM teachers
	
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []Group
	for rows.Next() {
		var group Group
		err := rows.Scan(&group.ID, &group.GroupName)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}

	return groups, nil
}

func getLessonsByGroupNameHandler(w http.ResponseWriter, r *http.Request) {
	// Подключение к базе данных
	db, err := connectDB()
	if err != nil {
		http.Error(w, "Ошибка подключения к базе данных", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Получение параметра group_name из строки запроса
	groupName := r.URL.Query().Get("group_name")
	if groupName == "" {
		http.Error(w, "Параметр group_name обязателен", http.StatusBadRequest)
		return
	}

	// Получение всех занятий для группы
	lessons, err := getLessonsByGroupName(db, groupName)
	if err != nil {
		http.Error(w, "Ошибка получения данных занятий", http.StatusInternalServerError)
		return
	}

	// Отправка результата в формате JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lessons)
}

func getLessonsByGroupName(db *sql.DB, groupName string) ([]Lesson, error) {
	var lessons []Lesson

	// SQL-запрос для получения всех занятий для группы
	query := `

		SELECT lessons.id, title, lesson_number, lesson_day
		FROM lessons
		JOIN groups ON lessons.group_id = groups.id JOIN teachers on lessons.teacher_id = teachers.id
		WHERE group_name = $1 or teacher_name = $1
	`

	rows, err := db.Query(query, groupName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// Обработка результатов запроса
	for rows.Next() {
		var lesson Lesson
		err := rows.Scan(&lesson.ID, &lesson.Title, &lesson.LessonNumber, &lesson.LessonDay)
		if err != nil {

			return nil, err
		}

		lessons = append(lessons, lesson)
	}
	return lessons, nil
}

func getStudentNameByInputDataHandler(w http.ResponseWriter, r *http.Request) {
	// Подключение к базе данных
	db, err := connectDB()
	if err != nil {
		http.Error(w, "Ошибка подключения к базе данных", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sPassword := r.URL.Query().Get("s_name")
	if sPassword == "" {
		http.Error(w, "Параметр имя обязателен", http.StatusBadRequest)
		return
	}
	sName := r.URL.Query().Get("s_password")
	if sName == "" {
		http.Error(w, "Параметр пароль обязателен", http.StatusBadRequest)
		return
	}

	res, err := getStudentByInputData(db, sName, sPassword)
	if err != nil {
		http.Error(w, "Ошибка получения данных занятий", http.StatusInternalServerError)
		return
	}

	// Отправка результата в формате JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func getStudentByInputData(db *sql.DB, sName string, sPassword string) (bool, error) {

	query := `
		SELECT * FROM students
		WHERE student_date = $1 and student_name = $2
	`

	rows, err := db.Query(query, sName, sPassword)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	if !rows.Next() {
		return false, nil // Нет строк, возвращаем false
	}

	// Обработка результатов запроса
	for rows.Next() {

		var student Student
		err := rows.Scan(&student.ID, &student.StudentDate, &student.GroupID, &student.StudentName)
		if err != nil {
			return false, err
		}

	}

	return true, nil
}

func getStudentMarksHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		http.Error(w, "Ошибка подключения к базе данных", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	sPassword := r.URL.Query().Get("s_password")
	sName := r.URL.Query().Get("s_name")

	marks, err := getStudentMarks(db, sPassword, sName)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(marks)
}

func getStudentMarks(db *sql.DB, sName string, sPassword string) ([]mark, error) {
	query := `SELECT mark_value, discipline, mark_date FROM mark JOIN mark_to_student ON mark_to_student.mark_id = mark.id JOIN students ON students.id = mark_to_student.student_id
		WHERE student_date = $1 and student_name = $2
		ORDER BY mark_date asc
	`

	rows, err := db.Query(query, sName, sPassword)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var marks []mark
	for rows.Next() {
		var mark mark
		err := rows.Scan(&mark.Value, &mark.Disciple, &mark.Date)
		if err != nil {
			return nil, err
		}
		marks = append(marks, mark)
	}
	return marks, nil
}

func getStudentHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		http.Error(w, "Ошибка подключения к базе данных", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sPassword := r.URL.Query().Get("s_password")
	if sPassword == "" {
		http.Error(w, "Параметр имя обязателен", http.StatusBadRequest)
		return
	}
	sName := r.URL.Query().Get("s_name")
	if sName == "" {
		http.Error(w, "Параметр пароль обязателен", http.StatusBadRequest)
		return
	}
	student, err := getStudent(db, sName, sPassword)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func getStudent(db *sql.DB, sName string, sPassword string) (string, error) {

	query := `
        SELECT group_name
        FROM students JOIN groups ON (group_id=groups.id)
        WHERE student_name = $1 AND student_date = $2
    `

	rows, err := db.Query(query, sName, sPassword)
	if err != nil {
		fmt.Println(err)
		return "no", err
	}
	defer rows.Close()

	var student string

	for rows.Next() {
		err := rows.Scan(&student)

		if err != nil {
			fmt.Println(err)

			return "no", err
		}
	}

	return student, nil
}

func getNewsHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		http.Error(w, "Ошибка подключения к базе данных", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	news, err := getNews(db)
	if err != nil {
		fmt.Println(err)

		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(news)

}

func getNews(db *sql.DB) ([]News, error) {
	query := `
		SELECT title, data, img, post_date FROM news
		ORDER BY post_date DESC 
`
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	var newsArray []News
	for rows.Next() {
		var news News
		err := rows.Scan(&news.Title, &news.Data, &news.Img, &news.PostDate)
		if err != nil {
			return nil, err
		}
		newsArray = append(newsArray, news)
	}
	return newsArray, nil
}

func handleRequest() {
	http.Handle("/Styles/", http.StripPrefix("/Styles/", http.FileServer(http.Dir("./Styles/"))))
	http.Handle("/Pages/", http.StripPrefix("/Pages/", http.FileServer(http.Dir("./Pages/"))))
	http.Handle("/Scripts/", http.StripPrefix("/Scripts/", http.FileServer(http.Dir("./Scripts/"))))
	http.Handle("/Images/", http.StripPrefix("/Images/", http.FileServer(http.Dir("./Images/"))))

	http.HandleFunc("/", index)
	http.HandleFunc("/api/groups", getGroupsHandler)

	http.HandleFunc("/api/lessons_by_group_name", getLessonsByGroupNameHandler)
	http.HandleFunc("/api/student_by_input_data", getStudentNameByInputDataHandler)
	http.HandleFunc("/api/student_data", getStudentHandler)
	http.HandleFunc("/api/student_marks", getStudentMarksHandler)
	http.HandleFunc("/api/news", getNewsHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}

func main() {
	db, err := connectDB()
	if err != nil {
		fmt.Println("Не удалось подключиться к базе данных. Завершение работы...")
		return
	}
	defer db.Close()

	handleRequest()
}
