package config

import (
	"fmt"

	"github.com/muhardiansyah15/go-rest-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:h@rd1@tcp(localhost:3306)/go_rest_api"))
	if err != nil {
		fmt.Println("Connection to database failed")
	}

	db.AutoMigrate(&models.User{})

	DB = db
}
