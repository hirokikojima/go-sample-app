package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:name`
	Email    string `json:email`
	Password string `json:password`
}

func (user *User)CreateUser(db *gorm.DB) {
	db.Create(&user)
}

func (u *User) FindUsers(db *gorm.DB) []User {
	users := []User{}
	db.Find(&users)
	return users
}

func (u *User)FindUser(db *gorm.DB) User {
	var user User
	db.Where("id = ?", u.ID).First(&user)
	return user
}

func (u *User) FindUserByEmail(db *gorm.DB) User {
	var user User
	db.Where("email = ?", u.Email).First(&user)
	return user
}