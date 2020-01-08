package routes

import (
	"golang-todo-mongo/handlers"
	"github.com/labstack/echo"
)

func InitiateBasicRoute(e *echo.Echo, groupName string){
	g := e.Group(groupName)
	g.GET("", handlers.Hello)
	g.POST("/signup", handlers.Signup)

}