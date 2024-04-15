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

type GetOvertimeSignOffSettingReq struct {
	DepartmentID uint `form:"departmentId" binding:"required"`
	VacationID   uint `form:"vacationId"`
}

type GetOvertimeSignOffSettingRes struct {
	*types.OvertimeSignOffSetting
}

type PutBatchOvertimeSignOffSettingReq struct {
	DepartmentID   uint
	VacationID     uint
	SignOffSetting []*types.OvertimeSignOffSetting `json:"signOffSetting"`
}

type PutBatchOvertimeSignOffSettingRes struct{}

type GetCheckInSignOffSettingReq struct{}

type GetCheckInSignOffSettingRes struct {
	*types.CheckInSignOffSetting
}

type PutBatchCheckInSignOffSettingReq struct {
	SignOffSetting []*types.CheckInSignOffSetting `json:"signOffSetting"`
}

type PutBatchCheckInSignOffSettingRes struct{}
