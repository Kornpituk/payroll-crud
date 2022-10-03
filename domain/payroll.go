package domain

import (
	"time"

	"gorm.io/gorm"
)

type Payroll struct {
	Id        int    `json:"hello_id"`
	Message   string `json:"message"`
	Name      string `json:"hello_name"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
