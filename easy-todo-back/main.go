package main

import (
	"easy-todo-back/dbConnection"
	"easy-todo-back/schema"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	db, sqldb, err := dbConnection.ConnectToDb()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer dbConnection.DisconnectToDb(sqldb)

	toDoDriver := schema.NewToDoDriver(db)

	r := gin.Default()

	r.POST("/api/v1/addnote", func(ctx *gin.Context) {
		val := ctx.Query("val")
		if err := toDoDriver.CreateToDo(val); err != nil {
			ctx.JSON(500, gin.H{"message": err})
		}
		ctx.JSON(200, gin.H{"message": "added!"})
	})

	r.GET("api/v1/alltodos", func(ctx *gin.Context) {
		todos, err := toDoDriver.ReadToDos()
		if err != nil {
			ctx.JSON(500, gin.H{"message": err})
		}
		ctx.JSON(200, todos)
	})

	r.GET("/api/v1/singletodo", func(ctx *gin.Context) {
		key, err := strconv.ParseInt(ctx.Query("key"), 10, 64)
		if err != nil {
			ctx.JSON(500, gin.H{"message": err})
		}
		todo, err := toDoDriver.ReadSingleTodoById(key)
		if err != nil {
			ctx.JSON(500, gin.H{"message": err})
		}
		ctx.JSON(200, todo)
	})

	r.PUT("/api/v1/rewrite", func(ctx *gin.Context) {
		key, err := strconv.ParseInt(ctx.Query("key"), 10, 64)
		if err != nil {
			ctx.JSON(500, gin.H{"message": err})
		}
		newVal := ctx.Query("newVal")
		toDoDriver.UpdateTodoValById(key, newVal)
		if err != nil {
			ctx.JSON(500, gin.H{"message": err})
		}
		ctx.JSON(200, gin.H{"message": "updated!"})
	})

	r.PUT("/api/v1/check", func(ctx *gin.Context) {
		key, err := strconv.ParseInt(ctx.Query("key"), 10, 64)
		if err != nil {
			ctx.JSON(500, gin.H{"message": "error"})
		}
		if err := toDoDriver.ToggleTodoById(key); err != nil {
			ctx.JSON(500, gin.H{"message": err})
		}
		ctx.JSON(200, gin.H{"message": "checked!"})
	})

	r.DELETE("/api/v1/delete", func(ctx *gin.Context) {
		key, err := strconv.ParseInt(ctx.Query("key"), 10, 64)
		if err != nil {
			ctx.JSON(500, gin.H{"message": "error"})
		}
		if err := toDoDriver.DeleteTodoById(key); err != nil {
			ctx.JSON(500, gin.H{"message": err})
		}
		ctx.JSON(200, gin.H{"message": "deleted!"})
	})

	r.Run()
}
