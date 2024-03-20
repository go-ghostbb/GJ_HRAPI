package types

import (
	"gorm.io/gorm"
	"time"
)

type CheckInData struct {
	gorm.Model
	CheckInDateTime time.Time `gorm:"comment:日期時間" json:"check_in_date_time"`
	EmployeeCode    string    `gorm:"size:50;comment:員工編號" json:"employee_code"`
}

func (c *CheckInData) TableName() string {
	return "check_in_data"
}
