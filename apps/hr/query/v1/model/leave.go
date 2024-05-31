package model

import (
	"hrapi/internal/types"
	"hrapi/internal/utils/response/page"
)

type GetLeaveCorrectReq struct {
	EmployeeID uint   `json:"-"`
	Year       string `form:"year" binding:"required"`
}

type GetLeaveCorrectRes struct {
	types.LeaveCorrect
}

type GetLeaveCorrectByKeywordReq struct {
	Keyword string `form:"keyword"`
	Year    string `form:"year" binding:"required"`
}

type GetLeaveCorrectByKeywordRes struct {
	page.Model[types.LeaveCorrect]
}
