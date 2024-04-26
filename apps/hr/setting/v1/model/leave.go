package model

import (
	"hrapi/internal/types"
	"hrapi/internal/utils/response/page"
)

type GetByKeywordLeaveReq struct {
	Keyword string `json:"keyword" form:"keyword"`
	Status  string `json:"status" form:"status"`
}

type GetByKeywordLeaveRes struct {
	page.Model[types.Leave]
}

type GetByIDLeaveReq struct {
	ID uint `json:"id" form:"id"`
}

type GetByIDLeaveRes struct {
	*types.Leave
}

type PostLeaveReq struct {
	*types.Leave
}

type PostLeaveRes struct{}

type PutLeaveReq struct {
	ID uint
	*types.Leave
}

type PutLeaveRes struct{}

type DeleteLeaveReq struct {
	ID uint
}

type DeleteLeaveRes struct{}

type SetStatusLeaveReq struct {
	ID     uint
	Status bool `json:"status"`
}

type SetStatusLeaveRes struct{}

type GetLeaveGroupReq struct {
	LeaveID uint
}

type GetLeaveGroupRes struct {
	*types.LeaveGroup
}

type PostLeaveGroupReq struct {
	*types.LeaveGroup
}

type PostLeaveGroupRes struct{}

type DeleteLeaveGroupReq struct {
	ID uint
}

type DeleteLeaveGroupRes struct{}

type SetLeaveGroupNameReq struct {
	ID   uint
	Name string `json:"name"`
}

type SetLeaveGroupNameRes struct{}

type SetLeaveGroupEmployeeReq struct {
	ID         uint
	EmployeeID []uint `json:"employeeId"`
}

type SetLeaveGroupEmployeeRes struct{}

type SetLeaveGroupCondReq struct {
	ID        uint
	Condition []*types.LeaveGroupCondition `json:"condition"`
}

type SetLeaveGroupCondRes struct{}

type ResetEmployeeAvailableReq struct {
	LeaveID    uint   `json:"leaveId"`
	EmployeeID []uint `json:"employeeId"`
	Year       int    `json:"year"`
}
