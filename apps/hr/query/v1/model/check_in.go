package model

import "hrapi/internal/types"

type QueryStatusReq struct {
	EmployeeID     uint
	DateRangeStart string `form:"dateRangeStart" binding:"required"`
	DateRangeEnd   string `form:"dateRangeEnd" binding:"required"`
	Abnormal       bool   `form:"abnormal"`
}

type QueryStatusRes struct {
	*types.CheckInStatus
}
