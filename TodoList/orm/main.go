package main

import (
	"gorm.io/gorm"
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

}
