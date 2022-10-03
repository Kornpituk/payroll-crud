package controller

import (
	"payroll/internal/controller/payroll"
	"payroll/internal/usecase"
)

type Service struct {
	Payroll *payroll.Service
}

func NewService(usecases *usecase.Service) *Service {
	return &Service{
		Payroll: payroll.NewService(usecases.Payroll),
	}
}
