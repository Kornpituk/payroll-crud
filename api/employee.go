package api

type EmployeeRequest struct {
	EmployeeId       string `json:"employeeId"`
	PayrollId        string `json:"payrollId"`
	LeaveId          string	`json:"leaveId"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	Gender           string `json:"gender"`
	Age              int    `json:"age"`
	ContactAddress   string `json:"contactAddress"`
	EmployeeEmail    string `json:"employeeEmail"`
	EmployeePassword string `json:"employeePassword"`
}

type EmployeeResponse struct {
	EmployeeId       string `gorm:"employeeId"`
	PayrollId        string `json:"payrollId"`
	LeaveId          string	`json:"leaveId"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	Gender           string `json:"gender"`
	Age              int    `json:"age"`
	ContactAddress   string `json:"contactAddress"`
	EmployeeEmail    string `json:"employeeEmail"`
	EmployeePassword string `json:"employeePassword"`
}
