package handlers

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/hirokikojima/go-sample-app/models"
	"github.com/hirokikojima/go-sample-app/utilities/authorizer"
)

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