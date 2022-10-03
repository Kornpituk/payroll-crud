package payroll

import (
	"payroll/internal/usecase/payroll"

)

type Service struct {
	payroll  *payroll.Service
}

func NewService(payroll *payroll.Service) *Service {
	return &Service{
		payroll: payroll,
	}
}
