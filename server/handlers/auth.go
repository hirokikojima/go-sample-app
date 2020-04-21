package handlers

import (
    "net/http"
	"github.com/labstack/echo"
	"github.com/hirokikojima/go-sample-app/models"
)

func Signup(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	if user.Email == "" || user.Password == "" {
		return &echo.HTTPError{
			Code: http.StatusBadRequest,
			Message: "invalid email or password",
		}
	}

	models.CreateUser(user)

	return c.JSON(http.StatusOK, user)
}

func Login(c echo.Context) error {
	user := new(models.User)
	return c.JSON(http.StatusOK, user)
}