package models

import (
	"github.com/jinzhu/gorm"
	"github.com/hirokikojima/go-sample-app/utilities"
)

func init() {
	db := utilities.ConnectDatabase()
	db.AutoMigrate(&User{})
}

type User struct {
	gorm.Model
	Name     string `json:name`
	Email    string `json:email`
	Password string `json:password`
}

func CreateUser(user *User) {
	db := utilities.ConnectDatabase()
	defer db.Close()
	db.Create(&user)
}

func FindUser(u *User) User {
	var user User
	db.Where(u).First(&user)
	return user
}