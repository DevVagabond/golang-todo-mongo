package main

import (
	"golang-todo-mongo/DB"
	"golang-todo-mongo/routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	DB.InitDb()
	routes.InitiateBasicRoute(e,"/api/v1/public")
	routes.InitiateSecureRoute(e,"/api/v1/secure")

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
