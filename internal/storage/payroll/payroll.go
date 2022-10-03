package payroll

import (
	"payroll/api"
	"payroll/internal/model"

	"gorm.io/gorm/clause"
)

func (s *Storage) GetPayroll() (*model.Payroll, error) {

	var model model.Payroll
	tx := s.db.First(&model)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return &model, nil
}

func (s *Storage) GetPayrolls(req api.PayrollRequest) ([]*model.Payroll, error) {

	var payroll []*model.Payroll

	// s.db.Preload("Employee").Preload("Leave").Preload("JobDepartment").Preload("Salary").Find(&payroll)

	tx := s.db.Model(&payroll).Preload(clause.Associations).Find(&payroll)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return payroll, nil
}

func (s *Storage) CreatePayroll(req api.PayrollRequest) (*model.Payroll, error) {

	var payroll *model.Payroll

	s.db.Preload("Employee").Preload("Leave").Preload("JobDepartment").Preload("Salary").Find(&payroll)

	model := model.Payroll{
		PayrollId:     req.PayrollId,
		Employee:      []model.Employee{},
		Leave:         []model.Leave{},
		JobDepartment: []model.JobDepartment{},
		Salary:        []model.Salary{},
		Date:          req.Date,
		Report:        req.Report,
		TotalAmount:   req.TotalAmount,
	}
	tx := s.db.Preload(clause.Associations).Create(&model)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &model, nil
}

func (s *Storage) GetJobPayrollById(id string) (*model.Payroll, error) {

	var payroll model.Payroll

	tx := s.db.Model(&payroll).Preload(clause.Associations).Where(&payroll.PayrollId,"=",id).Find(&payroll)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &payroll, nil
}

func (s *Storage) GetJobPayrollByIdEmp(id string) (*model.Payroll, error) {

	var payroll model.Payroll


	tx := s.db.Model(&payroll).Preload(clause.Associations).
	Select("payrolls.payroll_id","payrolls.date","employees.first_name",
	"employees.last_name","salaries.amount","salaries.bonus","payrolls.total_amount").
	Joins("INNER JOIN salaries on salaries.payroll_id = payrolls.payroll_id").
	Joins("INNER JOIN employees on employees.payroll_id = payrolls.payroll_id").
	Where(&payroll.PayrollId,"=",id).Find(&payroll)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &payroll, nil
}
