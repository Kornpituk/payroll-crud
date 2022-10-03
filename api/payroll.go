package api

import "time"

type PayrollRequest struct {
	PayrollId         string  `json:"payrollId"`
	Employee          string  `json:"employeeId"`
	EmployeeFirstName string  `json:"employeeFirstName"`
	EmployeeLastName  string  `json:"employeeLastName"`
	Leave             string  `json:"leaveId"`
	Salary            string  `json:"salaryId"`
	SalaryAmount      float64 `json:"salaryAmount"`
	SalaryBonus       float64 `json:"salaryBonus"`
	JobDepartment     string  `json:"jobDepartmentId"`
	Date              time.Time
	Report            string `json:"report"`
	TotalAmount       string `json:"totalAmount"`
}

type PayrollResponse struct {
	PayrollId         string `json:"payrollId"`
	Employee          string `json:"employeeId"`
	EmployeeFirstName string `json:"employeeFirstName"`
	EmployeeLastName  string `json:"employeeLastName"`
	Leave             string `json:"leaveId"`
	Salary            string `json:"salaryId"`
	SalaryAmount      float64 `json:"salaryAmount"`
	SalaryBonus       float64 `json:"salaryBonus"`
	JobDepartment     string `json:"jobDepartmentId"`
	Date              time.Time
	Report            string `json:"report"`
	TotalAmount       string `json:"totalAmount"`
}
