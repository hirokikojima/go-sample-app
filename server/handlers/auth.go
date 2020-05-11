package handlers

import (
	"log"
    "net/http"
	"github.com/labstack/echo"
	"github.com/hirokikojima/go-sample-app/models"
	"github.com/hirokikojima/go-sample-app/utilities/authorizer"
	"github.com/hirokikojima/go-sample-app/utilities/database"
)

func Signup(c echo.Context) error {
	db := database.ConnectDatabase()
	defer db.Close()

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

	encrypted, err := authorizer.EncryptPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = *encrypted
	user.CreateUser(db)
	user.Password = ""

	return c.JSON(http.StatusOK, user)
}

func Login(c echo.Context) error {
	db := database.ConnectDatabase()
	defer db.Close()

	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	u := user.FindUserByEmail(db)
	log.Printf("%v", u)
	log.Printf("%v", user)
	if !authorizer.ComparePassword(u.Password, user.Password) {
		return &echo.HTTPError{
            Code:    http.StatusUnauthorized,
            Message: "invalid email or password",
        }
	}

    return c.JSON(http.StatusOK, user)
}