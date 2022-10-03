package model

import "payroll/api"

type Salary struct {
	SalaryId      string          `gorm:"primaryKey;"`
	JobDepartment []JobDepartment `gorm:"foreignKey:SalaryId"`
	PayrollId     string
	Amount        float64
	Annual        float64
	Bonus         float64
}

func ToSalariesDomain(salaries []*Salary) []*api.SalaryResponse {

	results := make([]*api.SalaryResponse, 0, len(salaries))

	for _, v := range salaries {

		data := ToSalaryDomain(v)

		results = append(results, data)
	}

	return results
}

func ToSalaryDomain(salary *Salary) *api.SalaryResponse {

	results := &api.SalaryResponse{
		SalaryId:      salary.SalaryId,
		JobDepartment: salary.JobDepartment[0].JobId,
		PayrollId:     salary.PayrollId,
		Amount:        salary.Amount,
		Annual:        salary.Annual,
		Bonus:         salary.Bonus,
	}
	return results
}
