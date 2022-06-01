package database

import (
	"github.com/mashnoor/basic-api-golang/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connection() error {
	dbDsn := config.DBConfiguration()

	db, err := gorm.Open(postgres.Open(dbDsn), &gorm.Config{})

	if err != nil {
		return err
	}
	err = db.AutoMigrate(migrationModels...)
	if err != nil {
		log.Fatalln("Migration error")
	}

	DB = db

	return nil

}

func GetDB() *gorm.DB {
	return DB
}
