package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	ID          int
	Title       string
	IsCompleted bool
	CreatedAt   time.Time
}

func CreateTables(db *sql.DB) error {
	_, err := db.Exec(`
CREATE TABLE IF NOT EXISTS todos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT,
    is_completed INTEGER,
    created_at TEXT
)`)
	return err
}

func AddNewTodo(db *sql.DB, title string) error {
	_, err := db.Exec("INSERT INTO todos (title, is_completed, created_at) VALUES (?, ?, datetime('now'))", title, 0)
	if err != nil {
		return err
	}
	return nil
}

func ListTodos(db *sql.DB) ([]*Todo, error) {
	rows, err := db.Query("SELECT  * FROM todos ORDER  BY id desc ")
	if err != nil {
		return []*Todo{}, err
	}

	defer func() {
		_ = rows.Close()
	}()

	// ???????
	todos := []*Todo{}

	for rows.Next() {
		var id int
		var title string
		var isCompleted int
		var createsAtStr string

		if err := rows.Scan(&id, &title, &isCompleted, &createsAtStr); err != nil {
			return []*Todo{}, err
		}

		createdAt, err := time.Parse("2006-01-02 15:04:05", createsAtStr)
		if err != nil {
			return []*Todo{}, err
		}

		todos = append(todos, &Todo{
			ID:          id,
			Title:       title,
			IsCompleted: isCompleted == 1,
			CreatedAt:   createdAt,
		})

	}
	return todos, err
}

func CompleteTodo(db *sql.DB, id int) error {
	_, err := db.Exec("UPDATE todos SET is_completed=1 WHERE id = ?", id)
	return err
}

func main() {
	// connect to the sqlite database
	db, err := sql.Open("sqlite3", "file:database.sqlite3")
	if err != nil {
		panic(fmt.Sprintf("not able to connect to database: %s", err.Error()))
	}

	// Create the required tables
	if err := CreateTables(db); err != nil {
		panic(err)
	}

	// adding new todo
	http.HandleFunc("/new-todo", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != "POST" {
			http.Redirect(writer, request, "/", http.StatusPermanentRedirect)
			return
		}

		title := request.FormValue("title")

		if err := AddNewTodo(db, title); err != nil {
			http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
	})

	// done todo
	http.HandleFunc("/done-todo", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != "POST" {
			http.Redirect(writer, request, "/", http.StatusPermanentRedirect)
			return
		}

		idStr := request.FormValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(writer, "id must ne a number", http.StatusBadRequest)
			return
		}

		if err := CompleteTodo(db, id); err != nil {
			http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)

	})

	// Listing all todos
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

	if err := http.ListenAndServe(":9090", nil); err != nil {
		panic(err.Error())
	}

}
