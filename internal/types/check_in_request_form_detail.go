package types

import (
	"gorm.io/gorm"
	"hrapi/internal/types/enum"
	"hrapi/internal/utils/driver"
)

type CheckInRequestFormDetail struct {
	gorm.Model
	CheckInRequestFormID uint                `gorm:"not null;comment:補打卡單ID" json:"checkInRequestFormId"`
	CheckInRequestForm   *CheckInRequestForm `gorm:"foreignKey:CheckInRequestFormID;comment:補打卡單" json:"checkInRequestForm"`
	CheckInType          enum.CheckInType    `gorm:"comment:補打卡類型" json:"checkInType"`
	Date                 driver.Date         `gorm:"type:date;not null;comment:補打卡日期" json:"date"`
	Time                 driver.Time         `gorm:"type:time(0);not null;comment:補打卡時間" json:"time"`
	Remark               string              `gorm:"comment:補打卡原因" json:"remark"`
}

func (c *CheckInRequestFormDetail) TableName() string {
	return "check_in_request_form_detail"
}
