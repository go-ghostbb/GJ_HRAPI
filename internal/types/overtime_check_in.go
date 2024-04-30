package types

import (
	"gorm.io/gorm"
	"hrapi/internal/types/enum"
	"hrapi/internal/utils/driver"
)

type OvertimeCheckIn struct {
	gorm.Model
	EmployeeID  uint             `gorm:"comment:員工id" json:"employeeId"`
	Employee    Employee         `gorm:"foreignKey:EmployeeID" json:"employee"`
	Date        driver.Date      `gorm:"type:date;comment:日期" json:"date"`
	Time        driver.Time      `gorm:"type:time(0);comment:時間" json:"time"`
	CheckInType enum.CheckInType `gorm:"comment:加班上班或加班下班" json:"checkInType"`
}

func (o *OvertimeCheckIn) TableName() string {
	return "overtime_check_in"
}
