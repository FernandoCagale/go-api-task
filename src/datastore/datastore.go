package datastore

import (
	"fmt"

	"github.com/FernandoCagale/go-api-task/src/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	dbUser     = "test"
	dbPassword = "test"
	dbName     = "test"
	dbPort     = "5432"
)

func Connection() string {
	return fmt.Sprintf("host=api-postgres user=%s password=%s dbname=%s port=%s sslmode=disable", dbUser, dbPassword, dbName, dbPort)
}

func ConnectDB() (db *gorm.DB, err error) {
	db, err = gorm.Open("postgres", Connection())
	if err != nil {
		return
	}

	db.LogMode(false)

	db.AutoMigrate(&models.Task{})

	return
}
