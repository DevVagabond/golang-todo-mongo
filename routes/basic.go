package routes

import (
	"github.com/labstack/echo"
	"golang-todo-mongo/handlers"
)

func InitiateBasicRoute(e *echo.Echo, groupName string) {
	g := e.Group(groupName)
	g.POST("/signup", handlers.Signup)
	g.POST("/login", handlers.Login)

}
