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

	rec := db.ReadSingleTodoById(dbInstance, 2633542081492844029)
	log.Print(rec)
}
