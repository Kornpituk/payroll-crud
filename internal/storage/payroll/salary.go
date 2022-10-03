package payroll

import (
	"payroll/api"
	"payroll/internal/model"

	"gorm.io/gorm/clause"
)

func (s *Storage) GetSalaries(req api.SalaryRequest) ([]*model.Salary, error) {

	var salary []*model.Salary

	s.db.Preload("JobDepartment").Find(&salary)

	tx := s.db.Model(&salary).Preload(clause.Associations).Find(&salary)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return salary, nil
}

func (s *Storage) CreateSalary(req api.SalaryRequest) (*model.Salary, error) {

	model := model.Salary{
		SalaryId:      req.SalaryId,
		JobDepartment: []model.JobDepartment{},
		PayrollId:     req.PayrollId,
		Amount:        req.Amount,
		Annual:        req.Annual,
		Bonus:         req.Bonus,
	}

	tx := s.db.Create(&model).Preload(clause.Associations)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &model, nil
}

func (s *Storage) GetSalaryById(id string) (*model.Salary, error) {

	var salary model.Salary

	tx := s.db.Preload("JobDepartment").Where(&salary.SalaryId, "=", id).Find(&salary)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &salary, nil
}
