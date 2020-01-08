package middlewares

import (
	"github.com/labstack/echo"
	"net/http"
)

func CheckCookie(next echo.HandlerFunc)  echo.HandlerFunc{
	return func(c echo.Context) error {
		cookie,err := c.Cookie("sessionId")
		if err!= nil {
			return err
		}
		if cookie.Value == "sessionId" {
			return next(c)
		}
		return c.String(http.StatusUnauthorized,"unauthorized cookie");
	}
}