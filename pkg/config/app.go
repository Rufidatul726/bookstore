package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB is a global variable that represents the connection to the database
var DB *gorm.DB

// ConnectDB is a function to connect the database
func ConnectDB() {
	d, err := gorm.Open("mysql", "aurchey:aurchey@/bookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	DB = d
}

// GetDB is a function to get the connection to the database
func GetDB() *gorm.DB {
	return DB
}

// CloseDB is a function to close the connection to the database
func CloseDB() {
	DB.Close()
}

// Path: pkg\config\app.go
