package models

import (
	"github.com/jinzhu/gorm"
)

type Log struct {
	gorm.Model
	ServiceID uint   `json:service_id`
	Title     string `json:title`
	Body      string `json:body`
}

func (log *Log) CreateService(db *gorm.DB) {
	db.Create(&log)
}