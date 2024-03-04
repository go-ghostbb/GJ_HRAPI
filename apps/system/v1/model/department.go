package model

import "hrapi/internal/types"

type GetByKeywordDepartmentReq struct {
	Keyword string `json:"keyword"`
	Status  string `json:"status"`
}

type GetByKeywordDepartmentRes struct {
	*types.Department `json:",inline"`
}

type GetByIDDepartmentReq struct {
	ID uint `json:"id" form:"id"`
}

type GetByIDDepartmentRes struct {
	*types.Department `json:",inline"`
}

type PostDepartmentReq struct {
	*types.Department `json:",inline"`
}

type PostDepartmentRes struct{}

type PutDepartmentReq struct {
	ID                uint `form:"id"`
	*types.Department `json:",inline"`
}

type PutDepartmentRes struct{}

type DeleteDepartmentReq struct {
	ID uint `json:"id" form:"id"`
}

type DeleteDepartmentRes struct{}

type SetStatusDepartmentReq struct {
	ID     uint
	Status bool `json:"status" form:"status"`
}

type SetStatusDepartmentRes struct{}
