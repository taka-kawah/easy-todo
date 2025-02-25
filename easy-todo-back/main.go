package main

import (
	"easy-todo-back/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	dbInstance := db.ConnectToDb()
	sqldb, err := dbInstance.DB()
	if err != nil {
		panic("Error Getting SQLInstance")
	}
	defer db.DisconnectToDb(sqldb)

	r := gin.Default()

	r.POST("/api/v1/addnote", func(ctx *gin.Context) {
		val := ctx.Query("val")
		db.CreateToDo(dbInstance, val)
		ctx.JSON(200, gin.H{"message": "added!"})
	})

	r.GET("api/v1/alltodos", func(ctx *gin.Context) {
		todos := db.ReadToDo(dbInstance)
		ctx.JSON(200, todos)
	})

	r.GET("/api/v1/singletodo", func(ctx *gin.Context) {
		key, err := strconv.ParseInt(ctx.Query("key"), 10, 64)
		if err != nil {
			ctx.JSON(500, gin.H{"message": "error"})
		}
		todo := db.ReadSingleTodoById(dbInstance, key)
		ctx.JSON(200, todo)
	})

	r.PUT("/api/v1/rewrite", func(ctx *gin.Context) {
		key, err := strconv.ParseInt(ctx.Query("key"), 10, 64)
		if err != nil {
			ctx.JSON(500, gin.H{"message": "error"})
		}
		newVal := ctx.Query("newVal")
		db.UpDateTodoValById(dbInstance, key, newVal)
		ctx.JSON(200, gin.H{"message": "updated!"})
	})

	r.PUT("/api/v1/check", func(ctx *gin.Context) {
		key, err := strconv.ParseInt(ctx.Query("key"), 10, 64)
		if err != nil {
			ctx.JSON(500, gin.H{"message": "error"})
		}
		db.CheckTodoById(dbInstance, key)
		ctx.JSON(200, gin.H{"message": "checked!"})
	})

	r.DELETE("/api/v1/delete", func(ctx *gin.Context) {
		key, err := strconv.ParseInt(ctx.Query("key"), 10, 64)
		if err != nil {
			ctx.JSON(500, gin.H{"message": "error"})
		}
		db.DeleteTodoById(dbInstance, key)
		ctx.JSON(200, gin.H{"message": "deleted!"})
	})

	r.Run()
}
