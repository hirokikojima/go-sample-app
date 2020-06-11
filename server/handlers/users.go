package handlers

import (
	"net/http"
	"strconv"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/hirokikojima/go-sample-app/models"
	"github.com/hirokikojima/go-sample-app/responses"
	"github.com/hirokikojima/go-sample-app/utilities/database"
)

func GetUsers(c echo.Context) error {
	db := database.ConnectDatabase()
	defer db.Close()

	cc := NewCustomContext(c)

	user := models.User{}

	users := user.FindUsers(db)

	return cc.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	db := database.ConnectDatabase()
	defer db.Close()

	cc := NewCustomContext(c)

	userID, err := strconv.ParseUint(cc.Param("id"), 10, 64)
	if err != nil {
		return err
	}

	u := models.User{
		Model: gorm.Model{ID: uint(userID)},
	}

	user := u.FindUser(db)

	response := responses.NewUserResponse(user)

	return cc.JSON(http.StatusOK, response)
}