package model

import (
	"hrapi/internal/types"
	"hrapi/internal/utils/driver"
	"hrapi/internal/utils/response/page"
)

type QueryStatusReq struct {
	EmployeeID     uint
	DateRangeStart string `form:"dateRangeStart" binding:"required"`
	DateRangeEnd   string `form:"dateRangeEnd" binding:"required"`
	Abnormal       bool   `form:"abnormal"`
}

type QueryStatusRes struct {
	*types.CheckInStatus
}

type QueryStatusByKeywordReq struct {
	Keyword        string `form:"keyword"`
	DateRangeStart string `form:"dateRangeStart" binding:"required"`
	DateRangeEnd   string `form:"dateRangeEnd" binding:"required"`
	Abnormal       bool   `form:"abnormal"`
}

type QueryStatusByKeywordRes struct {
	page.Model[types.CheckInStatus]
}

type PutCheckInStatusReq struct {
	ID          uint
	WorkShiftID uint            `json:"workShiftId" binging:"required"`
	Work        driver.Datetime `json:"work"`
	OffWork     driver.Datetime `json:"offWork"`
}
