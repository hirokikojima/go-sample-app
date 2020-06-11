package handlers

import (
	"net/http"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/hirokikojima/go-sample-app/models"
	"github.com/hirokikojima/go-sample-app/responses"
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

	token, err := authorizer.GenerateSignedToken(user)
	if err != nil {
		return err
	}

    return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

func Login(c echo.Context) error {
	db := database.ConnectDatabase()
	defer db.Close()

	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	u := user.FindUserByEmail(db)
	if !authorizer.ComparePassword(u.Password, user.Password) {
		return &echo.HTTPError{
            Code:    http.StatusUnauthorized,
            Message: "invalid email or password",
        }
	}

	token, err := authorizer.GenerateSignedToken(&u)
	if err != nil {
		return err
	}

    return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

func Me(c echo.Context) error {
	db := database.ConnectDatabase()
	defer db.Close()

	claims := authorizer.GetClaims(c)

	u := models.User{
		Model: gorm.Model{ID: claims.UID},
	}

	user := u.FindUser(db)

	response := responses.NewUserResponse(user)

	return c.JSON(http.StatusOK, response)
}