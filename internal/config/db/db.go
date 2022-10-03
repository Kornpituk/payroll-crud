package db

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(conn string) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
