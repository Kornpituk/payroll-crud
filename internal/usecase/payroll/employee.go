package payroll

import (
	"net/http"
	"payroll/api"
	"payroll/internal/model"
)

func (s *Service) CreateEmployee(req api.EmployeeRequest) (*api.EmployeeResponse, *api.ErrorResponse) {

	var model *model.Employee

	model, err := s.storage.Payroll.CreateEmployee(req)

	if err != nil {
		return nil, &api.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	res := &api.EmployeeResponse{
		EmployeeId:       model.EmployeeId,
		PayrollId:        model.PayrollId,
		LeaveId:          model.LeaveId,
		FirstName:        model.FirstName,
		LastName:         model.LastName,
		Gender:           model.Gender,
		Age:              model.Age,
		ContactAddress:   model.ContactAddress,
		EmployeeEmail:    model.EmployeeEmail,
		EmployeePassword: model.EmployeePassword,
	}
	return res, nil
}

func (s *Service) GetEmployees(req api.EmployeeRequest) ([]*api.EmployeeResponse, *api.ErrorResponse) {

	models, err := s.storage.Payroll.GetEmployees(req)

	if err != nil {
		return nil, &api.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	results := model.ToEmployeesDomain(models)

	return results, nil
}

func (s *Service) GetEmployeeById(id string, sal api.SalaryRequest) (*api.EmployeeResponse, *api.ErrorResponse) {

	var model *model.Employee

	var err error

	model, err = s.storage.Payroll.GetEmployeeById(id)

	if err != nil {
		return nil, &api.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	result := &api.EmployeeResponse{
		EmployeeId:       model.EmployeeId,
		PayrollId:        model.PayrollId,
		LeaveId:          model.LeaveId,
		FirstName:        model.FirstName,
		LastName:         model.LastName,
		Gender:           model.Gender,
		Age:              model.Age,
		ContactAddress:   model.ContactAddress,
		EmployeeEmail:    model.EmployeeEmail,
		EmployeePassword: model.EmployeePassword,
	}
	return result, nil
}
