package models

import (
	"github.com/jinzhu/gorm"
)

type Service struct {
	gorm.Model
	UserID      uint   `json:user_id`
	Title       string `json:title`
	Body        string `json:body`
}

func (service *Service) CreateService(db *gorm.DB) {
	db.Create(&service)
}

func (service *Service) GetServices(db *gorm.DB) []Service {
	services := []Service{}
	db.Find(&services)
	return services
}