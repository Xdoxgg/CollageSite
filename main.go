package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"net/http"
)

type InfoPage struct {
	Name      string `json:"name"`
	InnerHTML string `json:"innerHTML"`
}

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

type Mark struct {
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

	groups, err := getGroups(db)
	if err != nil {
		http.Error(w, "Ошибка получения данных групп", http.StatusInternalServerError)
		return
	}

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

func getStudentMarks(db *sql.DB, sName string, sPassword string) ([]Mark, error) {
	query := `SELECT mark_value, discipline, mark_date FROM mark JOIN students ON students.id = mark.student_id
		WHERE student_date = $1 and student_name = $2
		ORDER BY mark_date asc
	`

	rows, err := db.Query(query, sName, sPassword)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var marks []Mark
	for rows.Next() {
		var mark Mark
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

func getInfoPageHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		http.Error(w, "Ошибка подключения к базе данных", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	data_id := r.URL.Query().Get("data_id")
	data, err := getInfoPage(db, data_id)
	if err != nil {
		fmt.Println(err)

		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)

}

func getInfoPage(db *sql.DB, id string) (*InfoPage, error) {
	query := `
	SELECT name, inner_html FROM info_pages WHERE id = $1
`

	rows, err := db.Query(query, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var ip InfoPage

	for rows.Next() {
		err := rows.Scan(&ip.Name, &ip.InnerHTML)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	return &ip, nil

}

func getGroupsOnlyHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		http.Error(w, "Ошибка подключения к базе данных", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	groups, err := getGroupsOnly(db)
	if err != nil {
		http.Error(w, "Ошибка получения данных групп", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groups)
}

func getGroupsOnly(db *sql.DB) ([]Group, error) {

	rows, err := db.Query(`
	SELECT id, group_name
	FROM groups


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

func getStudentsHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		http.Error(w, "Ошибка подключения к базе данных", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sPassword := r.URL.Query().Get("g_name")
	if sPassword == "" {
		http.Error(w, "Параметр имя обязателен", http.StatusBadRequest)
		return
	}

	student, err := getStudents(db, sPassword)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func getStudents(db *sql.DB, gName string) ([]Student, error) {

	query := `
       SELECT students.id, students.student_date, students.group_id, students.student_name
       FROM students JOIN groups ON students.group_id = groups.id WHERE group_name = $1
   `

	rows, err := db.Query(query, gName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	fmt.Println(gName)
	var students []Student
	for rows.Next() {
		var student Student
		//fmt.Println(student)
		err := rows.Scan(&student.ID, &student.StudentDate, &student.GroupID, &student.StudentName)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		students = append(students, student)
	}

	return students, nil
}

func setMarkHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		http.Error(w, "Ошибка подключения к базе данных", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	markValue := r.URL.Query().Get("mark_value")
	if markValue == "" {
		http.Error(w, "Параметр имя обязателен", http.StatusBadRequest)
		return
	}
	studentId := r.URL.Query().Get("student_id")
	if studentId == "" {
		http.Error(w, "Параметр имя обязателен", http.StatusBadRequest)
		return
	}
	discipline := r.URL.Query().Get("discipline")
	if discipline == "" {
		http.Error(w, "Параметр имя обязателен", http.StatusBadRequest)
		return
	}
	fmt.Println(markValue, studentId, discipline)
	if err := insertMark(db, markValue, studentId, discipline); err != nil {
		fmt.Println("err3")
		http.Error(w, "Ошибка при добавлении отметки", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nil)
}

func insertMark(db *sql.DB, markValue string, studentID string, discipline string) error {
	query := `
		INSERT INTO mark (mark_value, student_id, discipline, mark_date)
		VALUES ($1, $2, $3, NOW())
	`

	_, err := db.Exec(query, markValue, studentID, discipline) // Используем Exec для вставки
	if err != nil {
		fmt.Println(err) // Выводим ошибку в консоль
		return err       // Возвращаем ошибку
	}
	return nil
}
func handleRequest() {
	http.Handle("/Styles/", http.StripPrefix("/Styles/", http.FileServer(http.Dir("./Styles/"))))
	http.Handle("/Pages/", http.StripPrefix("/Pages/", http.FileServer(http.Dir("./Pages/"))))
	http.Handle("/Scripts/", http.StripPrefix("/Scripts/", http.FileServer(http.Dir("./Scripts/"))))
	http.Handle("/Images/", http.StripPrefix("/Images/", http.FileServer(http.Dir("./Images/"))))

	http.HandleFunc("/", index)
	http.HandleFunc("/api/groups", getGroupsHandler)
	http.HandleFunc("/api/groups_only", getGroupsOnlyHandler)

	http.HandleFunc("/api/lessons_by_group_name", getLessonsByGroupNameHandler)
	http.HandleFunc("/api/student_by_input_data", getStudentNameByInputDataHandler)
	http.HandleFunc("/api/student_data", getStudentHandler)
	http.HandleFunc("/api/student_marks", getStudentMarksHandler)
	http.HandleFunc("/api/news", getNewsHandler)
	http.HandleFunc("/api/info_data", getInfoPageHandler)
	http.HandleFunc("/api/students", getStudentsHandler)
	http.HandleFunc("/api/set_mark", setMarkHandler)

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
