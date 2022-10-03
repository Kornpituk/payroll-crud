package payroll

import (
	"payroll/api"
	"payroll/internal/model"
	"time"

	"gorm.io/gorm/clause"
)

func (s *Storage) CreateLeave(req api.LeaveRequest) (*model.Leave, error) {

	var leave model.Leave
	s.db.Preload("Employee").Find(&leave)

	leave = model.Leave{
		LeaveId:   req.LeaveId,
		Employee:  []model.Employee{},
		PayrollId: req.PayrollId,
		Date:      time.Time{},
		Status:    req.Status,
		Reason:    req.Reason,
	}

	tx := s.db.Create(&leave).Preload(clause.Associations).Association("Employee")
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &leave, nil

}

func (s *Storage) GetLeaves(req api.LeaveRequest) ([]*model.Leave, error) {

	var leaves []*model.Leave
	var employees []*model.Employee

	s.db.Model(&leaves).First(&leaves).Association("EmployeeId").Find(&employees)

	tx := s.db.Preload("Employee").Find(&leaves)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return leaves, nil
}

func (s *Storage) GetLeaveById(id string) (*model.Leave, error) {

	var leave model.Leave

	tx := s.db.Preload("Employee").Where(&leave.LeaveId, "=", id).Find(&leave)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &leave, nil
}
