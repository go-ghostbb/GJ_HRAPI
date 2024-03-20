package model

import "hrapi/internal/types"

type GetLeaveSignOffSettingReq struct {
	DepartmentID uint `form:"departmentId" binding:"required"`
	LeaveID      uint `form:"leaveId" binding:"required"`
}

type GetLeaveSignOffSettingRes struct {
	*types.LeaveSignOffSetting
}

type PutBatchLeaveSignOffSettingReq struct {
	DepartmentID   uint
	LeaveID        uint
	SignOffSetting []*types.LeaveSignOffSetting `json:"signOffSetting"`
}

type PutBatchLeaveSignOffSettingRes struct{}
