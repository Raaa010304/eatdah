package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("0.0.0.0:3306/eatdah"))
	if err != nil {
		panic(err)
	}
	database.AutoMigrate(&Menu{})
	database.AutoMigrate(&MenuFavorite{})

	DB = database
}
