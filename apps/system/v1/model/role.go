package model

import (
	"hrapi/internal/types"
	"hrapi/internal/utils/response/page"
)

type GetByKeywordRoleReq struct {
	Keyword string `json:"keyword" form:"keyword"`
	Status  string `json:"status" form:"status"`
}

type GetByKeywordRoleRes struct {
	page.Model[types.Role]
}

type GetByIDRoleReq struct {
	ID uint `form:"id"`
}

type GetByIDRoleRes struct {
	*types.Role
}

type PostRoleReq struct {
	*types.Role
}

type PostRoleRes struct{}

type PutRoleReq struct {
	ID uint `form:"id"`
	*types.Role
}

type PutRoleRes struct{}

type DeleteRoleReq struct {
	ID uint `form:"id"`
}

type DeleteRoleRes struct{}

type SetStatusRoleReq struct {
	ID     uint `json:"id" form:"id"`
	Status bool `json:"status" form:"status"`
}

type SetStatusRoleRes struct{}

type SetMenuRoleReq struct {
	ID     uint
	MenuId []uint `json:"menuId"`
}

type SetMenuRoleRes struct{}

type SetPermRoleReq struct {
	ID           uint
	PermissionId []uint `json:"permissionId"`
}

type SetPermRoleRes struct{}
