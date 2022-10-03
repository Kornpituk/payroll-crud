package storage

import (
	
	"payroll/internal/storage/payroll"

	"gorm.io/gorm"
)

type Service struct {
	Payroll  *payroll.Storage

}

func NewService(db *gorm.DB) *Service {
	return &Service{
		Payroll:  payroll.NewService(db),
	}
}


