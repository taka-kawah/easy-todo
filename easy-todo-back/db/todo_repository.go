package db

import (
	"easy-todo-back/schema"
	"log"

	"gorm.io/gorm"
)

func CreateToDo(db *gorm.DB, value string) {
	todo := schema.Todo{Id: schema.GetNewID(), Value: value, IsDone: false}
	res := db.Create(&todo)
	if res.Error != nil {
		panic("Error Creating New Todo Record")
	}
	log.Print("Success Creating New Todo Record!")
}

func ReadToDo(db *gorm.DB) []schema.Todo {
	var todos []schema.Todo
	res := db.Find(&todos)
	if err := res.Error; err != nil {
		panic("Error Reading Todo Records")
	}
	log.Print("Success Reading Todo Records")
	return todos
}
