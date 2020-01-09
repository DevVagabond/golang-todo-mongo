package handlers

import (
	"golang-todo-mongo/interfaces"
	"golang-todo-mongo/models"
	"encoding/json"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"time"
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
	hashedPwd,_ := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.MinCost)
	user.Password = string(hashedPwd);
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

func Login (c echo.Context) error {
	loginData := interfaces.LoginData{}
	user := interfaces.User{}
	defer c.Request().Body.Close()
	body,_ :=  ioutil.ReadAll(c.Request().Body)
	json.Unmarshal(body,&loginData)
	userObj := models.FindUser(loginData.Email)
	userObj.Decode(&user)
	if len(user.Email) <= 0 {
		return c.JSON(http.StatusNotFound,map[string]interface{}{
			"message" : "User not found",
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(loginData.Password))
	if(len(user.Email) <= 0 || err != nil) {
		return c.JSON(http.StatusBadRequest,map[string]interface{}{
			"message" : "Incorrect password or emailId",
		})
	}
	jwtkey := []byte("SECRET")
	expTime := time.Now().Add(5*time.Hour);
	claims := interfaces.Claims{
		Email:          user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt:expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	signedToken,err := token.SignedString(jwtkey)
	if(err != nil) {
		return  c.JSON(http.StatusBadRequest,map[string]string{
			"message" : "Something went wrong",
		})
	}

	return c.JSON(http.StatusOK,map[string]interface{}{
		"message" : "Found",
		"data" : map[string]interface{}{
			"_id": user.Id,
			"first_name" : user.FirstName,
			"last_name" : user.LastName,
			"email" : user.Email,
			"token" : signedToken,
		},
	})
}

