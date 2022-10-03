package api

type SalaryRequest struct {
	SalaryId      string  `json:"salaryId"`
	JobDepartment string  `json:"jobDepartmentId"`
	PayrollId       string  `json:"payrollId"`
	Amount        float64 `json:"amount"`
	Annual        float64 `json:"annual"`
	Bonus         float64 `json:"bonus"`
}

type SalaryResponse struct {
	SalaryId      string  `json:"salaryId"`
	JobDepartment string  `json:"jobDepartmentId"`
	PayrollId       string  `json:"payrollId"`
	Amount        float64 `json:"Amount"`
	Annual        float64 `json:"annual"`
	Bonus         float64 `json:"bonus"`
}
