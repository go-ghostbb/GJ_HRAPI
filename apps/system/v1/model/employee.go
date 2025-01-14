package model

import (
	"hrapi/internal/types"
	"hrapi/internal/types/enum"
	"hrapi/internal/utils/response/page"
)

type GetByKeywordEmployeeReq struct {
	Keyword          string                `form:"keyword" json:"keyword"`
	EmploymentStatus enum.EmploymentStatus `json:"employmentStatus" form:"employmentStatus"`
	DepartmentId     string                `form:"departmentId" json:"departmentId"`
}

type GetByKeywordEmployeeRes struct {
	page.Model[types.Employee]
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

type SetEmployeeRolesReq struct {
	EmployeeID uint
	RoleID     []uint `json:"roleId"`
}

type ResetPasswordReq struct {
	ID             uint
	Password       string `json:"password" v:"required#password cannot be empty|password#password must between 6, 18"`
	VerifyPassword string `json:"verifyPassword" v:"required#verify password cannot be empty|same:Password#the passwords entered twice do not match"`
}

type ResetPasswordRes struct{}

type SetEmploymentStatusReq struct {
	ID               uint
	EmploymentStatus enum.EmploymentStatus `json:"employmentStatus"`
}

type SetEmploymentStatusRes struct{}

type SetLoginStatusReq struct {
	ID     uint
	Status bool `json:"status"`
}

type SetLoginStatusRes struct{}

type GetByDateRangeCheckInStatusReq struct {
	DateRangeStart string `form:"dateRangeStart" binding:"required"`
	DateRangeEnd   string `form:"dateRangeEnd" binding:"required"`
	EmployeeID     uint   `form:"employeeId"`
	Abnormal       bool   `form:"abnormal"`
}

type GetByDateRangeCheckInStatusRes struct {
	*types.CheckInStatus
}

type GetByKeywordSalaryAddItemReq struct {
	Keyword    string `form:"keyword"`
	EmployeeID uint   `form:"employeeId"`
}

type GetByKeywordSalaryAddItemRes struct {
	ID           uint            `json:"ID"`
	Name         string          `json:"name"`
	SalaryType   enum.SalaryType `json:"salaryType"`
	IncomeTax    bool            `json:"incomeTax"`
	Benefits     bool            `json:"benefits"`
	Premiums     bool            `json:"premiums"`
	Amount       float32         `json:"amount"`
	UseCustom    bool            `json:"useCustom"`
	CustomAmount float32         `json:"customAmount"`
}

type SetAmountSalaryAddItemReq struct {
	EmployeeID      uint
	SalaryAddItemID uint    `json:"salaryAddItemId"`
	UseCustom       bool    `json:"useCustom"`
	CustomAmount    float32 `json:"customAmount"`
}
