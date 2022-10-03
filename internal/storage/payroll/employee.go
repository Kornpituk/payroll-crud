package payroll

import (
	"payroll/api"
	"payroll/internal/model"

	"gorm.io/gorm/clause"
)

func (s *Storage) CreateEmployee(req api.EmployeeRequest) (*model.Employee, error) {

	var employee model.Employee

	s.db.Preload("Payroll").Preload("Leave").Find(&employee)

	employee = model.Employee{
		EmployeeId:       req.EmployeeId,
		PayrollId:        req.PayrollId,
		LeaveId:          req.LeaveId,
		FirstName:        req.FirstName,
		LastName:         req.LastName,
		Gender:           req.Gender,
		Age:              req.Age,
		ContactAddress:   req.ContactAddress,
		EmployeeEmail:    req.EmployeeEmail,
		EmployeePassword: req.EmployeePassword,
	}

	tx := s.db.Create(&employee).Preload(clause.Associations)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &employee, nil
}

func (s *Storage) GetEmployees(req api.EmployeeRequest) ([]*model.Employee, error) {

	var employees []*model.Employee

	tx := s.db.Model(&employees).Find(&employees)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return employees, nil
}

func (s *Storage) GetEmployeeById(id string) (*model.Employee, error) {

	var employees model.Employee

	// s.db.Association("Employee").Find(&employees)

	tx := s.db.Model(&employees).Where(&employees.EmployeeId, "=", id).Find(&employees)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &employees, nil
}
