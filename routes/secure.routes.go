package routes

import (
	"github.com/labstack/echo"
)

func InitiateSecureRoute(e *echo.Echo, groupName string){
	g := e.Group(groupName)

	initiateToDoRoute(g);

	//g.GET("", handlers.Hello)
	//g.GET("/cats/:id", handlers.GetCat)
	//g.POST("/cats", handlers.PostCat)

}