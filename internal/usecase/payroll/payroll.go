package payroll

import (
	"net/http"
	"payroll/api"
	"payroll/internal/model"
	"time"
)

func (s *Service) GetPayrolls(req api.PayrollRequest) ([]*api.PayrollResponse, *api.ErrorResponse) {

	models, err := s.storage.Payroll.GetPayrolls(req)

	if err != nil {
		return nil, &api.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	results := model.ToPayrollsDomain(models)

	return results, nil
}

func (s *Service) CreatePayroll(req api.PayrollRequest) (*api.PayrollResponse, *api.ErrorResponse) {

	// model := make([]*model.Payroll, 0, 10)

	var model *model.Payroll

	model, err := s.storage.Payroll.CreatePayroll(req)

	if err != nil {
		return nil, &api.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	res := &api.PayrollResponse{
		PayrollId:     model.PayrollId,
		Employee:      model.Employee[0].EmployeeId,
		JobDepartment: model.JobDepartment[0].JobId,
		Salary:        model.Salary[0].SalaryId,
		Leave:         model.Leave[0].LeaveId,
		Date:          model.Date,
		Report:        model.Report,
		TotalAmount:   model.TotalAmount,
	}
	return res, nil
}

func (s *Service) GetPayrollById(id string) (*api.PayrollRequest, *api.ErrorResponse) {

	var model *model.Payroll

	var err error

	model, err = s.storage.Payroll.GetJobPayrollById(id)

	if err != nil {
		return nil, &api.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	result := &api.PayrollRequest{
		PayrollId:     id,
		Employee:      model.Employee[0].EmployeeId,
		Leave:         model.Leave[0].LeaveId,
		Salary:        model.Salary[0].SalaryId,
		JobDepartment: model.JobDepartment[0].JobId,
		Date:          time.Time{},
		Report:        model.Report,
		TotalAmount:   model.TotalAmount,
	}
	return result, nil
}

func (s *Service) GetPayrollByIdEmp(id string) (*api.PayrollRequest, *api.ErrorResponse) {

	var model *model.Payroll

	var err error

	model, err = s.storage.Payroll.GetJobPayrollByIdEmp(id)

	if err != nil {
		return nil, &api.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	result := &api.PayrollRequest{
		PayrollId:         model.PayrollId,
		JobDepartment:     model.JobDepartment[0].Name,
		EmployeeFirstName: model.Employee[0].FirstName,
		EmployeeLastName:  model.Employee[0].LastName,
		SalaryAmount:      model.Salary[0].Amount,
		SalaryBonus:       model.Salary[0].Bonus,
		Date:              time.Time{},
		Report:            model.Report,
		TotalAmount:       model.TotalAmount,
	}
	return result, nil
}
