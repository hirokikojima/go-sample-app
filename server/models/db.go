package models

import (
	"github.com/jinzhu/gorm"
	"github.com/hirokikojima/go-sample-app/utilities/database"
)

var db *gorm.DB

func init() {
	db := database.ConnectDatabase()
	defer db.Close()

	db.AutoMigrate(&User{})
}