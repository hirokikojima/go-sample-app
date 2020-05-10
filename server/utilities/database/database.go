package database

import (
	"github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func ConnectDatabase() *gorm.DB {
	err := godotenv.Load()
    if err != nil {
        panic(err)
    }
	DRIVER := os.Getenv("DATABASE_DRIVER")
	HOST := os.Getenv("DATABASE_HOST")
    USER := os.Getenv("DATABASE_USER")
    PASSWORD := os.Getenv("DATABASE_PASSWORD")
    DBNAME := os.Getenv("DATABASE_NAME")
    
    CONNECT := USER + ":" + PASSWORD + "@(" + HOST + ")/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(DRIVER, CONNECT)
    if err != nil {
        panic(err)
    }
    return db
}