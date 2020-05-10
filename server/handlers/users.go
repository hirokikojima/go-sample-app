package handlers

import (
    "net/http"
	"github.com/labstack/echo"
	"github.com/hirokikojima/go-sample-app/models"
	"github.com/hirokikojima/go-sample-app/utilities/database"
	"log"
)

func CreateUser(c echo.Context) error {
	log.Printf("create user")
	db := database.ConnectDatabase()
	defer db.Close()
	user := models.User{ Name: "Gopher" }
	db.Create(&user)
	return c.String(http.StatusOK, "こんにちは!")
}

func GetUser(c echo.Context) error {
	db := database.ConnectDatabase()
	defer db.Close()
	user := models.User{}
	db.First(&user, c.Param("id"))
	log.Printf(c.Param("id") + ":" + user.Name)
	return c.String(http.StatusOK, "こんにちは！")
}

func UpdateUser(c echo.Context) error {
	return c.String(http.StatusOK, "こんにちは!")
}

func DeleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "こんにちは!")
}