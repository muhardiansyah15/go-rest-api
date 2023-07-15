package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go_rest_api"))
	if err != nil {
		fmt.Println("Connection to database failed")
	}

	db.AutoMigrate(&User{})

	DB = db
}
