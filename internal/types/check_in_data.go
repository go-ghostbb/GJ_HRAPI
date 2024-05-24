package types

import (
	"gorm.io/gorm"
	"hrapi/internal/types/enum"
	"hrapi/internal/utils/driver"
)

type CheckInData struct {
	gorm.Model
	EmployeeID uint             `gorm:"comment:員工id" json:"employeeId"`
	Employee   *Employee        `gorm:"foreignKey:EmployeeID" json:"employee"`
	Datetime   driver.Datetime  `gorm:"type:datetime" json:"datetime"`
	Type       enum.CheckInType `gorm:"size:20;comment:類別" json:"type"`
}

func (c *CheckInData) TableName() string {
	return "check_in_data"
}
