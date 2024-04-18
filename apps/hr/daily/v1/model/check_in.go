package model

import (
	gbstr "ghostbb.io/gb/text/gb_str"
	"hrapi/internal/types"
	"hrapi/internal/types/enum"
	"hrapi/internal/utils/driver"
	"hrapi/internal/utils/response/page"
)

// -------------------form-------------------
type CheckInBasicForm struct {
	ID          uint                      `json:"id" copier:"ID"`
	EmployeeID  uint                      `json:"-" copier:"EmployeeID"`
	Detail      []*CheckInBasicFormDetail `json:"detail" copier:"Detail"`
	AttachArray []string                  `json:"attach" copier:"AttachArray"`
	Remark      string                    `json:"remark" copier:"Remark"`
}

type CheckInBasicFormDetail struct {
	CheckInType enum.CheckInType `gorm:"comment:補打卡類型" json:"checkInType" copier:"CheckInType"`
	Date        driver.Date      `gorm:"type:date;not null;comment:補打卡日期" json:"date" copier:"Date"`
	Time        driver.Time      `gorm:"type:time(0);not null;comment:補打卡時間" json:"time" copier:"Time"`
	Remark      string           `gorm:"comment:補打卡原因" json:"remark" copier:"Remark"`
}

func (c *CheckInBasicForm) Attach() string {
	return gbstr.Join(c.AttachArray, ",")
}

// ------------------------------------------

type GetCheckInFormByEmployeeReq struct {
	EmployeeID uint   `json:"employeeId"`
	StartDate  string `form:"startDate"`
	EndDate    string `form:"endDate"`
}

type GetCheckInFormByEmployeeRes struct {
	page.Model[types.CheckInRequestForm]
}

type SaveCheckInFormReq struct {
	CheckInBasicForm
}

type SaveCheckInFormRes struct {
	ID    uint   `json:"ID"`
	Order string `json:"order"`
}
