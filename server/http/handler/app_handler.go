package handler

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/hirokikojima/go-sample-app/models"
	"github.com/hirokikojima/go-sample-app/utilities/authorizer"
)

type AppHandler interface {
	AuthHandler
	UserHandler
	ServiceHandler
	LogHandler
}

type appHandler struct {
	AuthHandler
	UserHandler
	ServiceHandler
	LogHandler
}

func NewAppHandler(conn *gorm.DB) AppHandler {
	appHandler := &appHandler{}
	appHandler.AuthHandler = NewAuthHandler(conn)
	appHandler.UserHandler = NewUserHandler(conn)
	appHandler.ServiceHandler = NewServiceHandler(conn)
	appHandler.LogHandler = NewLogHandler(conn)
	return appHandler
}

type CustomContext struct {
	echo.Context
}

func NewCustomContext(c echo.Context) CustomContext {
	return CustomContext{c}
}

func (cc *CustomContext) CurrentUser(db *gorm.DB) models.User {
	claims := authorizer.GetClaims(cc)

	u := models.User{
		Model: gorm.Model{ID: claims.UID},
	}

	return u.FindUser(db)
}