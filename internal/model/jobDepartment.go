package model

import "payroll/api"

type JobDepartment struct {
	JobId         string `gorm:"primaryKey"`
	PayrollId     string
	SalaryId      string
	JobDepartment string
	Name          string
	Description   string
	SalaryRange   string
	Position      string
}

func ToJobDepartmentsDomain(jobDepartments []*JobDepartment) []*api.JobDepartmentResponse {

	var results []*api.JobDepartmentResponse

	for _, v := range jobDepartments {
		data := ToJobDepartmentDomain(v)
		results = append(results, data)
	}

	return results
}

func ToJobDepartmentDomain(jobDepartment *JobDepartment) *api.JobDepartmentResponse {

	results := &api.JobDepartmentResponse{
		JobId:         jobDepartment.JobId,
		PayrollId:     jobDepartment.PayrollId,
		SalaryId:      jobDepartment.SalaryId,
		JobDepartment: jobDepartment.JobDepartment,
		Name:          jobDepartment.Name,
		Description:   jobDepartment.Description,
		SalaryRange:   jobDepartment.SalaryRange,
		Position:      jobDepartment.Position,
	}
	return results
}
