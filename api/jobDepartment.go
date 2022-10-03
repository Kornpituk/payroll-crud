package api

type JobDepartmentRequest struct {
	JobId         string `json:"jobId"`
	PayrollId     string `json:"payrollId"`
	SalaryId      string `json:"SalaryId"`
	JobDepartment string `json:"jobDepartment"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	SalaryRange   string `json:"salaryRange"`
	Position      string `json:"position"`
}

type JobDepartmentResponse struct {
	JobId         string `json:"jobId"`
	PayrollId     string `json:"payrollId"`
	SalaryId      string `json:"SalaryId"`
	JobDepartment string `json:"jobDepartment"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	SalaryRange   string `json:"salaryRange"`
	Position      string `json:"position"`
}
