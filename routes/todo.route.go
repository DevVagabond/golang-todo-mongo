package routes

import (
	"github.com/labstack/echo"
	"golang-todo-mongo/handlers"
)

func initiateToDoRoute(secureGroup *echo.Group) {
	todoGroup := secureGroup.Group("/todo")
	todoGroup.GET("/list", handlers.List)
	todoGroup.POST("/add", handlers.CreateTodo)
	todoGroup.PUT("/status/:id", handlers.UpdateStatus)
	todoGroup.DELETE("/del/:id", handlers.RemoveTask)
}
