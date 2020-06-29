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

func (log *Log) CreateLog(db *gorm.DB) {
	db.Create(&log)
}

func (log *Log) UpdateLog(db *gorm.DB) {
	db.Save(&log)
}

func (log *Log) GetLogs(db *gorm.DB) []Log {
	logs := []Log{}
	db.Find(&logs)
	return logs
}