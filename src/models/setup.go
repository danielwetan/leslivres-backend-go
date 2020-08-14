package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/leslivres-go")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Book{})

	DB = database
}
