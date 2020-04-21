package models

import (
	"github.com/jinzhu/gorm"
	"github.com/hirokikojima/go-sample-app/utilities"
)

var db *gorm.DB

func init() {
	db := utilities.ConnectDatabase()
	defer db.Close()

	db.AutoMigrate(&User{})
}