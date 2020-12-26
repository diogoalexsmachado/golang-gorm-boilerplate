package db

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "newsfeed-go/models"
)

var db *gorm.DB

// instanciates DB and makes auto migrations
func Init() {
    db, err := gorm.Open(sqlite.Open("crm.db"), &gorm.Config{})

    if err != nil {
        fmt.Println(err.Error())
    }

    err = db.AutoMigrate(&models.Person{})

    if err != nil {
        fmt.Println(err.Error())
    }
}

//GetDB ...
func GetDB() *gorm.DB {
	return db
}