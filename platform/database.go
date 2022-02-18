package database_utils

import "gorm.io/gorm"

var DB *gorm.DB

func GetDatabase() (*gorm.DB) {
	return DB
}

func SetDatabase(db *gorm.DB) {
	DB = db
}