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
	ID          uint                `bson:"id" copier:"ID"`
	EmployeeID  uint                `json:"-"`
	Detail      []*CheckInBasicForm `json:"detail" copier:"Detail"`
	AttachArray []string            `json:"attach" copier:"AttachArray"`
}

type CheckInBasicFormDetail struct {
	CheckInType enum.CheckInType `gorm:"comment:補打卡類型" json:"checkInType"`
	Date        driver.Date      `gorm:"type:date;not null;comment:補打卡日期" json:"date"`
	Time        driver.Time      `gorm:"type:time(0);not null;comment:補打卡時間" json:"time"`
	Remark      string           `gorm:"comment:補打卡原因" json:"remark"`
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
