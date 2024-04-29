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
	Datetime    driver.Datetime  `gorm:"type:datetime;comment:日期時間" json:"datetime"`
	CheckInType enum.CheckInType `gorm:"comment:加班上班或加班下班" json:"checkInType"`
}

func (o *OvertimeCheckIn) TableName() string {
	return "overtime_check_in"
}
