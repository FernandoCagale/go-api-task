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
	return fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable", dbUser, dbPassword, dbName, dbPort)
}

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("postgres", connection())
	if err != nil {
		panic(err)
	}

	db.LogMode(false)

	db.AutoMigrate(&models.Task{})

	return db
}
