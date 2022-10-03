package model

import (
	"payroll/api"
	"time"
)

type Payroll struct {
	PayrollId     string          `gorm:"primaryKey"`
	Employee      []Employee      `gorm:"foreignKey:PayrollId"`
	Leave         []Leave         `gorm:"foreignKey:PayrollId"`
	JobDepartment []JobDepartment `gorm:"foreignKey:PayrollId"`
	Salary        []Salary        `gorm:"foreignKey:PayrollId"`
	Date          time.Time
	Report        string
	TotalAmount   string
}

func ToPayrollsDomain(payroll []*Payroll) []*api.PayrollResponse {



	results := make([]*api.PayrollResponse, 0 , len(payroll))

	for _, v := range payroll {

		data := ToPayrollDomain(v)
		
		
		results = append(results, data)

	}

	return results
}

func ToPayrollDomain(payroll *Payroll) *api.PayrollResponse {

	results := &api.PayrollResponse{
		PayrollId:     payroll.PayrollId,
		Employee:      payroll.Employee[0].EmployeeId,
		JobDepartment: payroll.JobDepartment[0].JobId,
		Salary:        payroll.Salary[0].SalaryId,
		Leave:         payroll.Leave[0].LeaveId,
		Date:          time.Time{},
		Report:        payroll.Report,
		TotalAmount:   payroll.TotalAmount,
	}
	return results
}
