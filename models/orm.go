package models

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func Connect() gorm.DB {
	url := os.Getenv("DB")
	if url == "" {
		url = "/tmp/gorm.db"
	}
	log.Println("Opening db ", url)
	db, err := gorm.Open("sqlite3", url)
	if err != nil {
		panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
	}
	if martini.Env == martini.Dev {
		//db.LogMode(true)
	}

	db.AutoMigrate(Media{})
	db.AutoMigrate(Group{})
	return db
}
