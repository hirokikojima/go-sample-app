package utilities

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func ConnectDatabase() *gorm.DB {
	DRIVER := os.Getenv("DATABASE_DRIVER")
	HOST := os.Getenv("DATABASE_HOST")
    USER := os.Getenv("DATABASE_USER")
    PASSWORD := os.Getenv("DATABASE_PASSWORD")
    DBNAME := os.Getenv("DATABASE_NAME")
    

    CONNECT := USER + ":" + PASSWORD + "@(" + HOST + ")/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local"
	log.Printf(CONNECT)
	db, err := gorm.Open(DRIVER, CONNECT)
    if err != nil {
        panic(err)
    }
    return db
}