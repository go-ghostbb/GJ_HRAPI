package model

import (
	gbstr "ghostbb.io/gb/text/gb_str"
	"hrapi/internal/types"
	"hrapi/internal/utils/driver"
	"hrapi/internal/utils/response/page"
	"time"
)

// -------------------form-------------------
type LeaveBasicForm struct {
	ID              uint     `json:"id" copier:"ID"`
	EmployeeID      uint     `json:"-"`
	ProxyEmployeeID uint     `json:"proxyEmployeeId" copier:"ProxyEmployeeID"`
	LeaveID         uint     `json:"leaveId" copier:"LeaveID"`
	DateArray       []string `json:"date" copier:"DateArray"`
	TimeArray       []string `json:"time" copier:"TimeArray"`
	AttachArray     []string `json:"attach" copier:"AttachArray"`
	Remark          string   `json:"remark" copier:"Remark"`
}

func (l *LeaveBasicForm) StartTime() driver.Time {
	result, err := time.Parse(time.TimeOnly, l.TimeArray[0])
	if err != nil {
		panic(err)
	}
	return driver.Time(result.Unix())
}

func (l *LeaveBasicForm) EndTime() driver.Time {
	result, err := time.Parse(time.TimeOnly, l.TimeArray[1])
	if err != nil {
		panic(err)
	}
	return driver.Time(result.Unix())
}

func (l *LeaveBasicForm) StartDate() driver.Date {
	result, err := time.Parse(time.DateOnly, l.DateArray[0])
	if err != nil {
		panic(err)
	}
	return driver.Date(result.Unix())
}

func (l *LeaveBasicForm) EndDate() driver.Date {
	result, err := time.Parse(time.DateOnly, l.DateArray[1])
	if err != nil {
		panic(err)
	}
	return driver.Date(result.Unix())
}

func (l *LeaveBasicForm) Attach() string {
	return gbstr.Join(l.AttachArray, ",")
}

// ------------------------------------------

type SaveLeaveFormReq struct {
	LeaveBasicForm
}

type SaveLeaveFormRes struct {
	ID    uint   `json:"ID"`
	Order string `json:"order"`
}

type GetLeaveFormByEmployeeReq struct {
	EmployeeID uint   `json:"employeeId"`
	StartDate  string `form:"startDate"`
	EndDate    string `form:"endDate"`
}

type GetLeaveFormByEmployeeRes struct {
	page.Model[types.LeaveRequestForm]
}

type DeleteLeaveFormReq struct {
	ID uint
}

type DeleteCheckInFormReq struct {
	ID uint
}
