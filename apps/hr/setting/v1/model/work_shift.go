package model

import (
	"hrapi/internal/types"
	"hrapi/internal/utils/response/page"
)

type GetByKeywordWorkShiftReq struct {
	Keyword string `json:"keyword" form:"keyword"`
	Status  string `json:"status" form:"status"`
}

type GetByKeywordWorkShiftRes struct {
	page.Model[types.WorkShift]
}

type GetByIDWorkShiftReq struct {
	ID uint `json:"id" form:"id"`
}

type GetByIDWorkShiftRes struct {
	*types.WorkShift
}

type PostWorkShiftReq struct {
	*types.WorkShift
}

type PostWorkShiftRes struct{}

type PutWorkShiftReq struct {
	ID uint
	*types.WorkShift
}

type PutWorkShiftRes struct{}

type DeleteWorkShiftReq struct {
	ID uint
}

type DeleteWorkShiftRes struct{}

type SetStatusWorkShiftReq struct {
	ID     uint
	Status bool `json:"status"`
}

type SetStatusWorkShiftRes struct{}

type GetByDateWorkScheduleReq struct {
	EmployeeID uint   `form:"employeeId" binding:"required"`
	Start      string `form:"start" binding:"required" v:"date-format:Y-m-d"`
	End        string `form:"end" binding:"required" v:"date-format:Y-m-d"`
}

type GetByDateWorkScheduleRes struct {
	*types.WorkSchedule
}

type PutBatchWorkScheduleReq struct {
	EmployeeID uint
	YearMonth  string                `json:"yearMonth"`
	Schedules  []*types.WorkSchedule `json:"schedules"`
}

type PutBatchWorkScheduleRes struct{}

type DeleteWorkScheduleReq struct {
	ID uint `form:"id" binding:"required"`
}

type QuickSettingWorkScheduleReq struct {
	Monday         []uint   `json:"mon"`
	Tuesday        []uint   `json:"tue"`
	Wednesday      []uint   `json:"wed"`
	Thursday       []uint   `json:"thu"`
	Friday         []uint   `json:"fri"`
	Saturday       []uint   `json:"sat"`
	Sunday         []uint   `json:"sun"`
	DateRange      []string `json:"dateRange" binding:"required"`
	EmployeeID     []uint   `json:"employeeId" binding:"required"`
	IgnoreVacation bool     `json:"ignoreVacation"`
}
