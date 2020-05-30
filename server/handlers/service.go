package handlers

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo"
	"github.com/hirokikojima/go-sample-app/models"
	"github.com/hirokikojima/go-sample-app/utilities/database"
)

func GetServices(c echo.Context) error {
	db := database.ConnectDatabase()
	defer db.Close()

	cc := NewCustomContext(c)

	service := models.Service{}

	if len(cc.Param("id")) > 0 {
		userID, err := strconv.ParseUint(cc.Param("id"), 10, 64)
		if err != nil {
			return err
		}
		service.UserID = uint(userID)
	}

	services := service.GetServices(db)

	return cc.JSON(http.StatusOK, services)
}

func CreateService(c echo.Context) error {
	db := database.ConnectDatabase()
	defer db.Close()

	cc := NewCustomContext(c)

	user := cc.CurrentUser(db)

	service := new(models.Service)
	if err := c.Bind(service); err != nil {
		return err
	}

	service.UserID = user.ID
	service.CreateService(db)

	return cc.JSON(http.StatusOK, service)
}