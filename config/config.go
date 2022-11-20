package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitDB(dsn string) {
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Can't connect database!")
	}

	log.Println("Connected to database")
}

func MirgrateDB() {
	DB.AutoMigrate()
}

func CloseDB() {
	dbSQL, _ := DB.DB()
	dbSQL.Close()
}
