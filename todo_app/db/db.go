package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
    //This port should be mysql's port
    d, err := gorm.Open("mysql","Bhuvnesh:@tcp(localhost:3306)/todo_app?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        panic(err)
    }
    db = d
}

func GetDB() *gorm.DB {
    if db == nil {
        log.Fatal("Database connection is not initialized")
    }
    return db
}


