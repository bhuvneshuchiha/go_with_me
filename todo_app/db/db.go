package db

import (
    "github.com/jinzhu/gorm"
)

var db *gorm.DB

func Connect() {
    d, err := gorm.Open("mysql","bhuvnesh:Bhuvnesh@123/todo_app/charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        panic(err)
    }
    db = d
}

func GetDB() *gorm.DB {
    return db
}
