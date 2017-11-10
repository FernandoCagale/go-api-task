package datastore

import (
	"github.com/FernandoCagale/go-api-task/src/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func New(connection string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", connection)
	if err != nil {
		return nil, err
	}

	db.LogMode(false)

	db.AutoMigrate(&models.Task{})

	return db, nil
}
