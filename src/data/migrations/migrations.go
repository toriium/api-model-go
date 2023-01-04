package main

import (
	"api/src/data/orm/tables"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	//DB_HOST := "127.0.0.1"
	DB_USER := "root"
	//DB_PORT := "3306"
	DB_NAME := "great_db"
	DB_PASSWORD := "123"
	dsn := DB_USER + ":" + DB_PASSWORD + "@tcp(127.0.0.1:3306)/" + DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&tables.TblBooks{})
	if err != nil {
		log.Fatal(err)
	}
}