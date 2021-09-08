package config

import (
	"aysf/day6r1/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type configDB struct {
	DB_Username string
	DB_Pass     string
	DB_Host     string
	DB_Port     string
	DB_Name     string
}

func InitDB() {

	conf := configDB{
		DB_Username: "root",
		DB_Pass:     "0123",
		DB_Host:     "127.0.0.1",
		DB_Port:     "3306",
		DB_Name:     "mvc_test",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DB_Username, conf.DB_Pass, conf.DB_Host, conf.DB_Port, conf.DB_Name)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	InitMigration()
}

func InitMigration() {
	DB.AutoMigrate(&models.User{})
}
