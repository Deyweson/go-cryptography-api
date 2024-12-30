package database

import (
	"log"

	"github.com/deyweson/go-cryptography-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DatabaseInitializer() {
	connectionURL := "host=localhost user=postgres password=0512 dbname=crypto port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(connectionURL))
	if err != nil {
		log.Panic(err.Error())
	}
	DB.AutoMigrate(&models.User{})
}
