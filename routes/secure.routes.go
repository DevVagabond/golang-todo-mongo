package routes

import (
	"github.com/labstack/echo"
	"golang-todo-mongo/middlewares"
)

func InitiateSecureRoute(e *echo.Echo, groupName string){
	g := e.Group(groupName)
	g.Use(middlewares.CheckJWT)
	initiateToDoRoute(g);
}