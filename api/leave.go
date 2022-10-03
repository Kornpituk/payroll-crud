package api

import "time"

type LeaveRequest struct {
	LeaveId  string `json:"leaveId"`
	Employee string `json:"employeeId"`
	PayrollId  string `json:"payrollId"`
	Date     time.Time
	Status   string `json:"status"`
	Reason   string `json:"reason"`
}

type LeaveResponse struct {
	LeaveId  string `json:"leaveId"`
	Employee string `json:"employeeId"`
	PayrollId  string `json:"payrollId"`
	Date     time.Time
	Status   string `json:"status"`
	Reason   string `json:"reason"`
}
