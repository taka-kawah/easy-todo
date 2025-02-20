package db

import (
	"easy-todo-back/schema"
	"log"

	"gorm.io/gorm"
)

func CreateToDo(db *gorm.DB, value string) {
	todo := schema.Todo{Id: schema.GetNewID(), Value: value, IsDone: false}
	log.Print(&todo)
	res := db.Create(&todo)
	if res.Error != nil {
		panic("Error Creating New Todo Record")
	}
	log.Print("Succerr Creating New Todo Record!", res.RowsAffected)
}
