package payroll

import (
	"payroll/api"
	"payroll/internal/model"
)

func (s *Storage) GetJobDeapartments(req api.JobDepartmentRequest) ([]*model.JobDepartment, error) {

	var models []*model.JobDepartment

	tx := s.db.Model(&models).Find(&models)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return models, nil
}

func (s *Storage) CreateJobDeapartment(req api.JobDepartmentRequest) (*model.JobDepartment, error) {

	model := model.JobDepartment{
		JobId:         req.JobId,
		PayrollId:     req.PayrollId,
		SalaryId:      req.SalaryId,
		JobDepartment: req.JobDepartment,
		Name:          req.Name,
		Description:   req.Description,
		SalaryRange:   req.SalaryRange,
		Position:      req.Position,
	}

	tx := s.db.Create(&model)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &model, nil
}

func (s *Storage) GetJobDeapartmentById(id string) (*model.JobDepartment, error) {

	var jobDepartment model.JobDepartment

	tx := s.db.Model(&jobDepartment).Where(&jobDepartment.JobId,"=",id).Find(&jobDepartment)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &jobDepartment, nil
}


