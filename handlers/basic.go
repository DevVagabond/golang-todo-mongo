package handlers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"golang-todo-mongo/interfaces"
	"golang-todo-mongo/models"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	"time"
)

// Handler

func Signup(c echo.Context) error {
	user := interfaces.User{}
	defer c.Request().Body.Close()
	body, _ := ioutil.ReadAll(c.Request().Body)
	json.Unmarshal(body, &user)
	hashedPwd, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	user.Password = string(hashedPwd)
	res, err := models.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Database error",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success signup",
		"data":    res,
	})
}

func Login(c echo.Context) error {
	loginData := interfaces.LoginData{}
	user := interfaces.User{}
	defer c.Request().Body.Close()
	body, _ := ioutil.ReadAll(c.Request().Body)
	json.Unmarshal(body, &loginData)
	userObj := models.FindUser(loginData.Email)
	userObj.Decode(&user)
	if len(user.Email) <= 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "User not found",
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if len(user.Email) <= 0 || err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Incorrect password or emailId",
		})
	}
	jwtkey := []byte("SECRET")
	expTime := time.Now().Add(5 * time.Hour)

	//fmt.Print(">>>>>>>>>>>>>>>", user, "///////////", user.Id)

	claims := interfaces.Claims{
		Email:  user.Email,
		UserId: user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtkey)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Something went wrong",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Found",
		"data": map[string]interface{}{
			"_id":        user.Id,
			"first_name": user.FirstName,
			"last_name":  user.LastName,
			"email":      user.Email,
			"token":      signedToken,
		},
	})
}
