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

func main() {

}
