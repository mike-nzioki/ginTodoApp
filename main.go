package main

import (
	"github.com/gin-gonic/gin"
	"todo/handlers"
	"todo/repositories"
)

func main() {
	router := gin.Default()

	ConnectDatabase()
	todoRepository := repositories.NewTodoRepository(DB)
	todoHandler := handlers.NewTodoHandler(todoRepository)

	router.GET("/todos", todoHandler.GetTodos)
	router.POST("/todos", todoHandler.CreateTodo)
	router.GET("/todos/:id", todoHandler.GetTodoByID)
	router.PATCH("/todos/:id", todoHandler.UpdateTodo)
	router.DELETE("/todos/:id", todoHandler.DeleteTodo)

	router.Run(":8080")
}
