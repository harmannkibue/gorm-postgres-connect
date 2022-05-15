package db

import (
	"example/postgres-connect/pkg/common/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Init(dsn string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Failed to Open gorm:>", err)
	}

	err = db.AutoMigrate(&models.Book{})

	if err != nil {
		fmt.Println("Failed to run migrations")
		return nil
	}

	return db
}
