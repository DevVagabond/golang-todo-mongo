package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"golang-todo-mongo/interfaces"
	"net/http"
)

func CheckJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		haeder_token := c.Request().Header.Get("x-access-token")
		claims := &interfaces.Claims{}
		decoded, err := jwt.ParseWithClaims(haeder_token, claims, func(token *jwt.Token) (i interface{}, err error) {
			return []byte("SECRET"), nil
		})
		if err != nil || !decoded.Valid {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "Error",
			})
		}
		c.Request().Header.Set("email", claims.Email)
		c.Request().Header.Set("userId", claims.UserId.Hex())
		return next(c)
	}
}
