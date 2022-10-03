package payroll

import (
	"net/http"
	"payroll/api"
	"payroll/internal/model"
)

func (s *Service) GetJobDepartments(req api.JobDepartmentRequest) ([]*api.JobDepartmentResponse, *api.ErrorResponse) {

	models, err := s.storage.Payroll.GetJobDeapartments(req)

	if err != nil {
		return nil, &api.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	results := model.ToJobDepartmentsDomain(models)

	return results, nil
}

func (s *Service) CreateDepartment(req api.JobDepartmentRequest) (*api.JobDepartmentResponse, *api.ErrorResponse) {

	var model *model.JobDepartment

	model, err := s.storage.Payroll.CreateJobDeapartment(req)

	if err != nil {
		return nil, &api.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	res := &api.JobDepartmentResponse{
		JobId:         model.JobId,
		PayrollId:     model.PayrollId,
		SalaryId:      model.SalaryId,
		JobDepartment: model.JobDepartment,
		Name:          model.Name,
		Description:   model.Description,
		SalaryRange:   model.SalaryRange,
		Position:      model.Position,
	}
	return res, nil
}

func (s *Service) GetJobDepartmentById(id string) (*api.JobDepartmentResponse, *api.ErrorResponse) {

	var model *model.JobDepartment

	var err error

	model, err = s.storage.Payroll.GetJobDeapartmentById(id)

	if err != nil {
		return nil, &api.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	result := &api.JobDepartmentResponse{
		JobId:         model.JobId,
		PayrollId:     model.PayrollId,
		SalaryId:      model.SalaryId,
		JobDepartment: model.JobDepartment,
		Name:          model.Name,
		Description:   model.Description,
		SalaryRange:   model.SalaryRange,
		Position:      model.Position,
	}
	return result, nil
}
