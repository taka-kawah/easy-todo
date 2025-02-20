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

	db.CreateToDo(dbInstance, "ゴミ捨て")
}
