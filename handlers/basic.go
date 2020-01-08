package handlers

import (
	"golang-todo-mongo/interfaces"
	"golang-todo-mongo/models"
	"encoding/json"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
)

// Handler
func Hello(c echo.Context) error {
	v := c.Param("q")
	return c.JSON(http.StatusOK, map[string]string {
		"paramsss" : v,
	})
}


func Signup (c echo.Context) error {
	user := interfaces.User{}
	defer c.Request().Body.Close()
	body,_ := ioutil.ReadAll(c.Request().Body)
	json.Unmarshal(body,&user);
	res,err := models.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest,map[string]string{
			"message" : "Database error",
		})
	}
	return c.JSON(http.StatusOK,map[string]interface{}{
		"message" : "Success signup",
		"data" : res,
	})
}

