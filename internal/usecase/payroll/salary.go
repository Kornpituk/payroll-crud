package payroll

import (
	"net/http"
	"payroll/api"
	"payroll/internal/model"
)

func (s *Service) GetSalaries(req api.SalaryRequest) ([]*api.SalaryResponse, *api.ErrorResponse) {

	models, err := s.storage.Payroll.GetSalaries(req)

	if err != nil {
		return nil, &api.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	results := model.ToSalariesDomain(models)

	return results, nil
}

func (s *Service) CreateSalary(req api.SalaryRequest) (*api.SalaryResponse, *api.ErrorResponse) {

	var model *model.Salary

	model, err := s.storage.Payroll.CreateSalary(req)

	if err != nil {
		return nil, &api.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	res := &api.SalaryResponse{
		SalaryId:      model.SalaryId,
		JobDepartment: model.JobDepartment[0].JobId,
		PayrollId:     model.PayrollId,
		Amount:        model.Amount,
		Annual:        model.Annual,
		Bonus:         model.Bonus,
	}
	return res, nil
}

func (s *Service) GetSalaryById(id string) (*api.SalaryResponse, *api.ErrorResponse) {

	var model *model.Salary

	var err error

	model, err = s.storage.Payroll.GetSalaryById(id)

	if err != nil {
		return nil, &api.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	result := &api.SalaryResponse{
		SalaryId:      model.SalaryId,
		JobDepartment: model.JobDepartment[0].JobId,
		PayrollId:     model.PayrollId,
		Amount:        model.Amount,
		Annual:        model.Annual,
		Bonus:         model.Bonus,
	}
	return result, nil
}
