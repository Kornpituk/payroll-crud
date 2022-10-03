package payroll

import (
	"net/http"
	"payroll/api"
	"payroll/internal/model"
	"time"
)

func (s *Service) CreateLeave(req api.LeaveRequest) (*api.LeaveResponse, *api.ErrorResponse) {

	var model *model.Leave

	model, err := s.storage.Payroll.CreateLeave(req)

	if err != nil {
		return nil, &api.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	res := &api.LeaveResponse{
		LeaveId:  model.LeaveId,
		Employee: model.Employee[0].EmployeeId,
		Date:     time.Time{},
		Status:   model.Status,
		Reason:   model.Reason,
	}
	return res, nil
}

func (s *Service) GetLeaves(req api.LeaveRequest) ([]*api.LeaveResponse, *api.ErrorResponse) {

	models, err := s.storage.Payroll.GetLeaves(req)

	if err != nil {
		return nil, &api.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	results := model.ToLevesDomain(models)

	return results, nil
}

func (s *Service) GetLeaveById(id string) (*api.LeaveResponse, *api.ErrorResponse) {

	var model *model.Leave

	var err error

	model, err = s.storage.Payroll.GetLeaveById(id)

	if err != nil {
		return nil, &api.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	result := &api.LeaveResponse{
		LeaveId:  model.LeaveId,
		Employee: model.Employee[0].EmployeeId,
		PayrollId:  model.PayrollId,
		Date:     time.Time{},
		Status:   model.Status,
		Reason:   model.Reason,
	}
	return result, nil
}
