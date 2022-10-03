package payroll

import "gorm.io/gorm"
//p
type Storage struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Storage {
	return &Storage{
		db: db,
	}
}
