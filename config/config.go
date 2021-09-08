package config

import (
	"fmt"

	"github.com/aysf/gojwt/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	fmt.Println("executed 2")
	connectionString := "root:0123@tcp(127.0.0.1:3306)/latihan1?charset=utf8&parseTime=True&loc=Local"
	fmt.Println("executed 3")

	var err error

	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	InitialMigration()
}

func InitialMigration() {
	DB.AutoMigrate(&models.Book{})
	DB.AutoMigrate(&models.User{})
}
