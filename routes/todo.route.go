package routes
import (
	"golang-todo-mongo/handlers"
	"github.com/labstack/echo"
)

func initiateToDoRoute(secureGroup *echo.Group) {
	todoGroup := secureGroup.Group("/todo")
	todoGroup.GET("/list",handlers.List)
}