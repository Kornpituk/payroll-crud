package usecase

import (
	"payroll/internal/config"
	"payroll/internal/storage"
	"payroll/internal/usecase/payroll"

)

type Service struct {
	Payroll  *payroll.Service

}

func NewService(config *config.Config, storage *storage.Service) *Service {
	return &Service{
		Payroll: payroll.NewService(config, storage),
	}
}
