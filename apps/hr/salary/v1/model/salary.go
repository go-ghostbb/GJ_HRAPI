package model

import (
	"hrapi/internal/types"
	"hrapi/internal/utils/response/page"
)

type SalaryCalcReq struct {
	FounderEmployeeID uint
	EmployeeID        []uint `json:"employeeId"`
	DateRangeStart    string `form:"dateRangeStart" binding:"required"`
	DateRangeEnd      string `form:"dateRangeEnd" binding:"required"`
}

type SalaryCalcRes struct {
	*types.CalcSalary
}

type GetSalaryCalcReq struct{}

type GetSalaryCalcRes struct {
	page.Model[types.CalcSalary]
}

type DeleteSalaryCalcReq struct {
	ID uint
}

type GetByIDSalaryCalcReq struct {
	ID uint
}

type GetByIDSalaryCalcRes struct {
	*types.CalcSalary
}

type SaveEmployeeItemReq struct {
	CalcSalaryEmployee []*types.CalcSalaryEmployee `json:"calcSalaryEmployee"`
}
