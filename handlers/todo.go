package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang-todo-mongo/interfaces"
	"golang-todo-mongo/models"
	"io/ioutil"
	"net/http"
)

func List(c echo.Context) error {
	list := models.FindTodo(c.Request().Header.Get("userId"))
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    list,
	})
}
func CreateTodo(c echo.Context) error {
	todoData := interfaces.Todo{}
	defer c.Request().Body.Close()
	body, _ := ioutil.ReadAll(c.Request().Body)
	err := json.Unmarshal(body, &todoData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error",
		})
	}
	todoData.Completed = false
	todoData.UserId, _ = primitive.ObjectIDFromHex(c.Request().Header.Get("userId"))
	data, _ := models.CreateTodo(todoData)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    data,
	})
}

func UpdateStatus(c echo.Context) error {
	id := c.Param("id")
	var todoData interfaces.Todo
	defer c.Request().Body.Close()
	body, _ := ioutil.ReadAll(c.Request().Body)
	_ = json.Unmarshal(body, &todoData)
	fmt.Print(todoData)
	res, err := models.UpdateTodo(id, todoData)
	fmt.Print(">>>>>>>>>>>>  ", err)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Updated",
		"data":    res,
	})
}

func RemoveTask(c echo.Context) error {
	id := c.QueryParam("id")
	res, err := models.RemoveTodo(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Updated",
		"data":    res,
	})
}
