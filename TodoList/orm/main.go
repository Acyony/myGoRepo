package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

type Todo struct {
	ID          int `gorm:"primarykey"`
	Title       string
	IsCompleted bool
	CreatedAt   time.Time
}

func ListTodos(db *gorm.DB) ([]*Todo, error) {
	todos := []*Todo{}
	result := db.Find(&todos)
	return todos, result.Error
}

func CompleteTodo(db *gorm.DB, id int) error {
	result := db.Model(&Todo{}).Where("id=?", id).Update("is_completed", true)
	return result.Error
}

func AddNewTodo(db *gorm.DB, title string) error {
	todo := &Todo{
		Title:       title,
		IsCompleted: false,
		CreatedAt:   time.Now(),
	}

	result := db.Create(todo)

	return result.Error
}
func main() {
	//connect to the sqlite database
	db, err := gorm.Open(sqlite.Open("database_gorm.sqlite3"), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("not able to connect to database: %s", err.Error()))
	}

	// create the required tables
	if err := db.AutoMigrate(&Todo{}); err != nil {
		panic(fmt.Sprintf("not able to create a table: %s", err.Error()))
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		todos, err := ListTodos(db)
		if err != nil {
			http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		indexTmpl := template.Must(template.ParseFiles("template/index.gohtml"))
		writer.Header().Set("Content-Type", "text/html")

		_ = indexTmpl.Execute(writer, todos)
	})

	http.HandleFunc("/done-todo", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != "Post" {
			http.Redirect(writer, request, "/", http.StatusPermanentRedirect)
			return
		}

		idStr := request.FormValue("id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			http.Error(writer, "id must be a number", http.StatusBadRequest)
			return
		}

		if err := CompleteTodo(db, id); err != nil {
			http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
	})

}
