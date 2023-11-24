package database

import (
	"github.com/hacktiv8-fp-golang/final-project-04/internal/config"
	"github.com/hacktiv8-fp-golang/final-project-04/internal/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	err error
)

func StartDB() {
	config := config.GetDBConfig()

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	db.Debug().AutoMigrate(model.User{}, model.Category{}, model.Product{}, model.TransactionHistory{})
}

func GetDB() *gorm.DB {
	return db
}
