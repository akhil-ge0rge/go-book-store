package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const dsn = "akhil:akhil123@tcp(localhost:3306)/simplerest?charset=utf8mb4&parseTime=True&loc=Local"

func Connect() {

	d, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{},
	)

	if err != nil {
		panic(err)
	}
	DB = d

	log.Println("Database connected successfully")
}

func GetDB() *gorm.DB {
	return DB
}
