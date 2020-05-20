package handlers

import (
    "net/http"
	"github.com/labstack/echo"
	"github.com/hirokikojima/go-sample-app/models"
	"github.com/hirokikojima/go-sample-app/utilities/database"
)

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