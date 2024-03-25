package model

import "hrapi/internal/types"

type SalaryCalcReq struct {
	EmployeeID []uint `json:"employeeId"`
}

type SalaryCalcRes struct {
	*types.CalcSalary
}
