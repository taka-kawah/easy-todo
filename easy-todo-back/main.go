package main

import (
	"easy-todo-back/db"
)

func main() {
	dbInstance := db.ConnectToDb()
	sqldb, err := dbInstance.DB()
	if err != nil {
		panic("Error Getting SQLInstance")
	}
	defer db.DisconnectToDb(sqldb)

	targetId := 2633542081492844029

	db.DeleteTodoById(dbInstance, int64(targetId))
}
