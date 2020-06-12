package models

import (
	"github.com/jinzhu/gorm"
)

type Service struct {
	gorm.Model
	UserID      uint   `json:user_id`
	Title       string `json:title`
	Body        string `json:body`
	User        User   `json:user`
	Logs        []Log  `json:logs`
}

func (service *Service) CreateService(db *gorm.DB) {
	db.Create(&service)
}

func (service *Service) GetServices(db *gorm.DB) []Service {
	services := []Service{}
	db.Set("gorm:auto_preload", true).Find(&services)
	return services
}