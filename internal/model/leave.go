package model

import (
	"payroll/api"
	"time"
)

type Leave struct {
	LeaveId   string     `gorm:"primaryKey"`
	Employee  []Employee `gorm:"foreignKey:LeaveId"`
	PayrollId string
	Date      time.Time
	Status    string
	Reason    string
}

func ToLevesDomain(leave []*Leave) []*api.LeaveResponse {

	var results []*api.LeaveResponse

	for _, v := range leave {
		data := ToLeaveDomain(v)
		results = append(results, data)
	}

	return results
}

func ToLeaveDomain(leave *Leave) *api.LeaveResponse {

	results := &api.LeaveResponse{
		LeaveId:   leave.LeaveId,
		Employee:  leave.Employee[0].EmployeeId,
		PayrollId: leave.PayrollId,
		Date:      time.Time{},
		Status:    leave.Status,
		Reason:    leave.Reason,
	}
	return results
}
