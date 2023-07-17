package database

import (
	"chatapp/pkg/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func InitialiseDatabase() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", "user", "user", "localhost", "prise")
	open, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf(err.Error())
	}

	database, _ := open.DB()

	err = database.Ping()
	if err != nil {
		panic(err.Error())
	}
	err = open.AutoMigrate(&models.ProductPriceFees{})
	if err != nil {
		log.Fatalf("Error:%v", err)
	}

	if err != nil {
		log.Fatalf(err.Error())
	}
	db = open
	return db
}

func GetInstance() *gorm.DB {
	return db
}
