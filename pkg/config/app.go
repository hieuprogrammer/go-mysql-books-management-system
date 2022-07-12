package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	database, error := gorm.Open("mysql", "root:root@/book?charset=utf8&parseTime=True&loc=Local")
	if error != nil {
		panic(error)
	}

	db = database
}

func GetDb() *gorm.DB {
	return db
}
