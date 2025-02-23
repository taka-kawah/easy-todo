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
	log.Print("Success Reading Todo Records!")
	return todos
}

func ReadSingleTodoById(db *gorm.DB, id int64) schema.Todo {
	var todos []schema.Todo
	res := db.First(&todos, id)
	if err := res.Error; err != nil || len(todos) != 1 {
		panic(("Error Reading Single Todo Record"))
	}
	log.Print("Success Reading Single Todo Record!")
	return todos[0]
}

func UpDateTodoValById(db *gorm.DB, id int64, newVal string) {
	var todos []schema.Todo
	res := db.Model(&todos).Where("id = ?", id).Update("Value", newVal)
	if err := res.Error; err != nil {
		panic(("Error Updating Todo Record"))
	}
	log.Print("Success Updating Todo Record Value!")
}

func CheckTodoById(db *gorm.DB, id int64) {
	var todos []schema.Todo
	res := db.Model(&todos).Where("id = ?", id).Update("is_done", gorm.Expr("NOT is_done"))
	if err := res.Error; err != nil {
		panic(("Error Checking Todo Record"))
	}
	log.Print("Success Checking Todo Record Value!")
}
