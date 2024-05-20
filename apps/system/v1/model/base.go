package model

import "hrapi/internal/types"

type LoginReq struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type LoginRes struct {
	Token string `json:"token"`
}

type GetEmployeeInfoReq struct {
}

type GetEmployeeInfoRes struct {
	EmployeeId   uint              `json:"employeeId" copier:"ID"`
	DepartmentId uint              `json:"departmentId" copier:"DepartmentId"`
	Department   *types.Department `json:"department" copier:"Department"`
	Account      string            `json:"account"`
	RealName     string            `json:"realName" copier:"RealName"`
	Avatar       string            `json:"avatar" copier:"Avatar"`
	Roles        []*RoleInfo       `json:"roles" copier:"Roles"`
}

func (g *GetEmployeeInfoRes) LoginInformation(info *types.LoginInformation) {
	g.Account = info.Account
}

type RoleInfo struct {
	Code string `json:"code" copier:"Code"`
	Name string `json:"name" copier:"Name"`
}

type GetMenuReq struct{}

type GetMenuRes struct {
	ParentID  uint          `json:"-" copier:"ParentID"`
	Sort      int           `json:"-" copier:"Sort"`
	Path      string        `json:"path" copier:"Path"`
	Component string        `json:"component" copier:"Component"`
	Meta      types.Meta    `json:"meta" copier:"Meta"`
	Name      string        `json:"name" copier:"Name"`
	Redirect  string        `json:"redirect" copier:"Redirect"`
	Children  []*GetMenuRes `json:"children" copier:"-"`
}

type ChangePasswordReq struct {
	EmployeeID     uint   `json:"-"`
	OldPassword    string `json:"oldPassword" v:"required#old password cannot be empty"`
	NewPassword    string `json:"newPassword" v:"required#new password cannot be empty|password#password must between 6, 18"`
	VerifyPassword string `json:"verifyPassword" v:"required#verify password cannot be empty|same:NewPassword#the passwords entered twice do not match"`
}

type ChangePasswordRes struct{}

type CreateMasterKeyReq struct {
	Count int    `form:"count" bind:"required"`
	Unit  string `form:"unit" bind:"required"`
}

type CreateMasterKeyRes struct {
	Token string `json:"token"`
}
