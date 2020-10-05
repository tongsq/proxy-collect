package model

import "github.com/jinzhu/gorm"

var (
	DB *gorm.DB
)

func init() {
	DB = New()
}
