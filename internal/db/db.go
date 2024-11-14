package db

import (
	"campaign-service/util"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var store *gorm.DB

func Connect(c util.Config) (*gorm.DB, error) {
	var err error
	store, err = gorm.Open(postgres.Open(c.DBSource), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = Migrate()
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	return store, nil
}

func Migrate() error {
	err := store.AutoMigrate(
		&Campaign{},
		&CampaignUser{},
	)
	if err != nil {
		return err
	}
	return nil
}
