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
	recs := db.ReadToDo(dbInstance)
	for _, rec := range recs {
		log.Print(rec)
	}
}
