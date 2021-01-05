package mysql

import "github.com/jinzhu/gorm"
import _ "github.com/jinzhu/gorm/dialects/mysql"

var (
	db *gorm.DB
)

func DB() *gorm.DB {
	if db == nil {
		db = NewDB()
	}
	return db
}

func NewDB() (db *gorm.DB) {
	db, err := gorm.Open("mysql", "python:123456@(127.0.0.1:3306)/py?charset=utf8&loc=Local")
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	return db
}
