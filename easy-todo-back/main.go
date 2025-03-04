package main

import (
	"easy-todo-back/dbConnection"
	"easy-todo-back/middleware"
	"easy-todo-back/schema"
	"log"
	"net/http"
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

	var ud *schema.UserDriver
	toDoDriver := schema.NewToDoDriver(db)

	e := gin.Default()
	r := e.Group("/api/v1")
	r.POST("/SignUp", func(ctx *gin.Context) {
		email := ctx.Query("email")
		password := ctx.Query("password")

		var err error
		ud, err = schema.NewUserDriver(db, email, password)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
	})

	r.POST("/SignIn", func(ctx *gin.Context) {
		email := ctx.Query("email")
		middleware.LoginHandler(ctx)
		var err error
		ud, err = schema.FindUserDriverFromEmail(db, email)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
	})

	authorized := r.Group("/auth")
	authorized.Use(middleware.AuthMiddleware())

	authorized.POST("/addnote", func(ctx *gin.Context) {
		val := ctx.Query("val")
		if err := toDoDriver.CreateToDo(val, ud.ID); err != nil {
			ctx.JSON(500, gin.H{"message": err})
		}
		ctx.JSON(200, gin.H{"message": "added!"})
	})

	authorized.GET("/alltodos", func(ctx *gin.Context) {
		todos, err := toDoDriver.ReadToDos()
		if err != nil {
			ctx.JSON(500, gin.H{"message": err})
		}
		ctx.JSON(200, todos)
	})

	authorized.GET("/singletodo", func(ctx *gin.Context) {
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

	authorized.PUT("/rewrite", func(ctx *gin.Context) {
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

	authorized.PUT("/check", func(ctx *gin.Context) {
		key, err := strconv.ParseInt(ctx.Query("key"), 10, 64)
		if err != nil {
			ctx.JSON(500, gin.H{"message": "error"})
		}
		if err := toDoDriver.ToggleTodoById(key); err != nil {
			ctx.JSON(500, gin.H{"message": err})
		}
		ctx.JSON(200, gin.H{"message": "checked!"})
	})

	authorized.DELETE("/delete", func(ctx *gin.Context) {
		key, err := strconv.ParseInt(ctx.Query("key"), 10, 64)
		if err != nil {
			ctx.JSON(500, gin.H{"message": "error"})
		}
		if err := toDoDriver.DeleteTodoById(key); err != nil {
			ctx.JSON(500, gin.H{"message": err})
		}
		ctx.JSON(200, gin.H{"message": "deleted!"})
	})

	e.Run()
}
