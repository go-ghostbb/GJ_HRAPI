package model

import (
	"hrapi/internal/types"
	"hrapi/internal/utils/driver"
	"hrapi/internal/utils/response/page"
)

// -------------------form-------------------
type OvertimeBasicForm struct {
	ID         uint        `json:"id" copier:"ID"`
	EmployeeID uint        `json:"-" copier:"EmployeeID"`
	Date       driver.Date `json:"date" copier:"Date"`
	StartTime  driver.Time `json:"startTime" copier:"StartTime"`
	EndTime    driver.Time `json:"endTime" copier:"EndTime"`
	Remark     string      `json:"remark" copier:"Remark"`
}

// ------------------------------------------

type GetOvertimeFormByEmployeeReq struct {
	EmployeeID uint   `json:"employeeId"`
	StartDate  string `form:"startDate"`
	EndDate    string `form:"endDate"`
}

type GetOvertimeFormByEmployeeRes struct {
	page.Model[types.OvertimeRequestForm]
}

type SaveOvertimeFormReq struct {
	OvertimeBasicForm
}

type SaveOvertimeFormRes struct {
	ID    uint   `json:"ID"`
	Order string `json:"order"`
}

type DeleteOvertimeFormReq struct {
	ID uint
}
