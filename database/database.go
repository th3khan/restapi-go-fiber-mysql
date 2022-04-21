package database

import (
	"log"
	"os"

	"github.com/th3khan/restapi-go-fiber-mysql/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectionDb() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect the database.!\n", err.Error())
		os.Exit(2)
	}

	log.Println("Database connected successfully!")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations...")
	// TODO: Add migrations
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	Database = DbInstance{Db: db}
}
