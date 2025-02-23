package main

import (
	"easy-todo-back/db"
	"log"
)

func main() {
	dbInstance := db.ConnectToDb()
	sqldb, err := dbInstance.DB()
	if err != nil {
		panic("Error Getting SQLInstance")
	}
	defer db.DisconnectToDb(sqldb)

	targetId := 2633542081492844029

	db.CheckTodoById(dbInstance, int64(targetId))
	rec := db.ReadSingleTodoById(dbInstance, int64(targetId))
	log.Print(rec)
}
