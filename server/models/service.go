package models

import (
	"github.com/jinzhu/gorm"
)

type Service struct {
	gorm.Model
	UserID      uint   `json:user_id`
	Name        string `json:name`
	Description string `json:description`
}

func (service *Service) CreateService(db *gorm.DB) {
	db.Create(&service)
}