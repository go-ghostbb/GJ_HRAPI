package model

import (
	"hrapi/internal/types"
	"hrapi/internal/types/enum"
)

type GetByKeywordEmployeeReq struct {
	Keyword          string                `form:"keyword" json:"keyword"`
	EmploymentStatus enum.EmploymentStatus `json:"employmentStatus" form:"employmentStatus"`
	DepartmentId     string                `form:"departmentId" json:"departmentId"`
}

type GetByKeywordEmployeeRes struct {
	*types.Employee
}

type GetByIDEmployeeReq struct {
	ID uint `form:"id" json:"id"`
}

type GetByIDEmployeeRes struct {
	*types.Employee
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
