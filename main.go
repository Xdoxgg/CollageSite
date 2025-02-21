﻿package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"encoding/json"
	_ "github.com/lib/pq" // Используем "_" для импорта драйвера pq
)

type Group struct {
	ID        int    `json:"id"`
	GroupName string `json:"group_name"`
}

type Student struct {
	ID          int    `json:"id"`
	GroupID     int    `json:"group_id"`
	StudentName string `json:"student_name"`
}

type Lesson struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	LessonNumber int    `json:"lesson_number"`
	LessonDay 	 int 	`json:"lesson_day"`
	LessonDate   int    `json:"lesson_date"`
	Place		 int 	`json:"place"`
	GroupID 	 int 	 `json:"group_id"`
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

func handleRequest() {
	http.Handle("/Styles/", http.StripPrefix("/Styles/", http.FileServer(http.Dir("./Styles/"))))
	http.Handle("/Pages/", http.StripPrefix("/Pages/", http.FileServer(http.Dir("./Pages/"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/api/groups", getGroupsHandler) // Новый API-эндпоинт
	
	http.HandleFunc("/api/lessons_by_group_name", getLessonsByGroupNameHandler)
	// http.HandleFunc("/api/test", testHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}



func getGroups(db *sql.DB) ([]Group, error) {
	rows, err := db.Query(`SELECT id, group_name FROM groups`)
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
		JOIN groups ON lessons.group_id = groups.id
		WHERE group_name = $1
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

func main() {
	db, err := connectDB()
	if err != nil {
		fmt.Println("Не удалось подключиться к базе данных. Завершение работы...")
		return
	}
	defer db.Close()



	handleRequest()
}