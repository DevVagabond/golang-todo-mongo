package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

func List(c echo.Context) error{
	return c.String(http.StatusOK,"ok")
};