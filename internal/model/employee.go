package model

import (
	"payroll/api"
)

type Employee struct {
	EmployeeId       string `gorm:"primaryKey"`
	PayrollId        string
	LeaveId          string
	FirstName        string
	LastName         string
	Gender           string
	Age              int
	ContactAddress   string
	EmployeeEmail    string
	EmployeePassword string
}

func ToEmployeesDomain(employees []*Employee) []*api.EmployeeResponse {

	var results []*api.EmployeeResponse

	for _, v := range employees {
		data := ToEmployeeDomain(v)
		results = append(results, data)
	}

	return results
}

func ToEmployeeDomain(employee *Employee) *api.EmployeeResponse {

	results := &api.EmployeeResponse{
		EmployeeId:       employee.EmployeeId,
		PayrollId:        employee.PayrollId,
		LeaveId:          employee.LeaveId,
		FirstName:        employee.FirstName,
		LastName:         employee.LastName,
		Gender:           employee.Gender,
		Age:              employee.Age,
		ContactAddress:   employee.ContactAddress,
		EmployeeEmail:    employee.EmployeeEmail,
		EmployeePassword: employee.EmployeePassword,
	}
	return results
}
