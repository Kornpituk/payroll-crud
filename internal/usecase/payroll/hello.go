package payroll

import "payroll/api"

func (s *Service) Hello() (*api.HelloResponse, *api.ErrorResponse) {
	res := &api.HelloResponse{
		Message: "Hello Auto Migration!",
	}
	return res, nil
}
