package model

import (
	"hrapi/internal/types"
	"hrapi/internal/utils/driver"
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

type PutLeaveCorrectReq struct {
	ID        uint
	Effective driver.Date `json:"effective"`
	Expired   driver.Date `json:"expired"`
	Available float64     `json:"available"`
}
