package model

import "hrapi/internal/types"

type GetLeaveCorrectReq struct {
	EmployeeID uint   `json:"-"`
	Year       string `form:"year" binding:"required"`
}

type GetLeaveCorrectRes struct {
	types.LeaveCorrect
}
