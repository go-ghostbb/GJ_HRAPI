package model

import "hrapi/internal/types"

type GetByKeywordEmployeeReq struct {
	Keyword string `form:"keyword" json:"keyword"`
}

type GetByKeywordEmployeeRes struct {
	Employees []*types.Employee
}

type GetByIDEmployeeReq struct {
	ID uint `form:"id" json:"id"`
}

type GetByIDEmployeeRes struct {
	Employee *types.Employee
}

type PostEmployeeReq struct {
	*types.Employee `json:",inline"`
}

type PostEmployeeRes struct {
	Account  string
	Password string
}

type PutEmployeeReq struct {
	ID       uint `form:"id"`
	Employee *types.Employee
}

type PutEmployeeRes struct {
}

type DeleteEmployeeReq struct {
	ID uint `form:"id"`
}

type DeleteEmployeeRes struct {
}
