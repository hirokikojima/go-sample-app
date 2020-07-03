package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewDBConnection() *gorm.DB {
	return getMySqlConnection()
}

func getMySqlConnection() *gorm.DB {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		Current.Database.User,
		Current.Database.Password,
		Current.Database.Host,
		Current.Database.Port,
		Current.Database.Database,
	)

	conn, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	err = conn.DB().Ping()
	if err != nil {
		panic(err)
	}

	return conn
}