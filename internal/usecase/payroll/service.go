package payroll

import (
	"payroll/internal/config"
	"payroll/internal/storage"
)

type Service struct {
	config  *config.Config
	storage *storage.Service
}

func NewService(config *config.Config, storage *storage.Service) *Service {
	return &Service{
		config:  config,
		storage: storage,
	}
}
