package model

import (
	gbstr "ghostbb.io/gb/text/gb_str"
	"hrapi/internal/types"
	"hrapi/internal/types/enum"
	"hrapi/internal/utils/response/page"
	"time"
)

// -------------------form-------------------
type LeaveForm struct {
	LeaveBasicForm
	LeaveFormInfo
	FormOther
	FormLeaveCount
}

type FormOther struct {
	DepartmentName        string          `json:"departmentName" copier:"DepartmentName"`
	ProxyEmployeeRealName string          `json:"proxyEmployeeRealName" copier:"ProxyEmployeeRealName"`
	ProxyDepartmentName   string          `json:"proxyDepartmentName" copier:"ProxyDepartmentName"`
	LeaveName             string          `json:"leaveName" copier:"LeaveName"`
	LeaveMinimum          uint            `json:"leaveMinimum" copier:"LeaveMinimum"`
	SignStatus            enum.SignStatus `json:"signStatus" copier:"SignStatus"`
}

type FormLeaveCount struct {
	LeaveMinuteCount float32 `json:"leaveMinuteCount" copier:"LeaveMinuteCount"`
	LeaveHourCount   float32 `json:"leaveHourCount" copier:"LeaveHourCount"`
	LeaveDayCount    float32 `json:"leaveDayCount" copier:"LeaveDayCount"`
}

type LeaveFormInfo struct {
	CreatedAt         time.Time `json:"createdAt" copier:"CreatedAt"`
	Order             string    `json:"order" copier:"Order"`
	DepartmentID      uint      `json:"departmentId" copier:"DepartmentID"`
	ProxyDepartmentID uint      `json:"proxyDepartmentId" copier:"ProxyDepartmentID"`
}

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

func (l *LeaveBasicForm) StartTime() time.Time {
	result, err := time.Parse(time.TimeOnly, l.TimeArray[0])
	if err != nil {
		panic(err)
	}
	return result
}

func (l *LeaveBasicForm) EndTime() time.Time {
	result, err := time.Parse(time.TimeOnly, l.TimeArray[1])
	if err != nil {
		panic(err)
	}
	return result
}

func (l *LeaveBasicForm) StartDate() time.Time {
	result, err := time.Parse(time.DateOnly, l.DateArray[0])
	if err != nil {
		panic(err)
	}
	return result
}

func (l *LeaveBasicForm) EndDate() time.Time {
	result, err := time.Parse(time.DateOnly, l.DateArray[1])
	if err != nil {
		panic(err)
	}
	return result
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
	StartDate  string `form:"startDate" binding:"required"`
	EndDate    string `form:"endDate" binding:"required"`
}

type GetLeaveFormByEmployeeRes struct {
	page.Model[types.LeaveRequestForm]
}

type DeleteLeaveFormReq struct {
	ID uint
}
