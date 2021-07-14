package services

import (
	"github.com/ocleyson/hero-app-api/models"
	"github.com/ocleyson/hero-app-api/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	DB_HOST := utils.GetEnvVar("DB_HOST")
	DB_USER := utils.GetEnvVar("DB_USER")
	DB_PASSWORD := utils.GetEnvVar("DB_PASSWORD")
	DB_NAME := utils.GetEnvVar("DB_NAME")
	DB_PORT := utils.GetEnvVar("DB_PORT")

	dsn := "host=" + DB_HOST + " user=" + DB_USER + " password=" + DB_PASSWORD + " dbname=" + DB_NAME + " port=" + DB_PORT + " sslmode=disable"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.Hero{})

	DB = database
}
