package types

import (
	"gorm.io/gorm"
	"hrapi/internal/utils/driver"
)

type LeaveDefer struct {
	gorm.Model
	EmployeeID uint        `gorm:"default:0;not null;comment:員工ID"`
	Employee   *Employee   `gorm:"foreignKey:EmployeeID;comment:員工"`
	LeaveID    uint        `gorm:"default:0;not null;comment:假別ID"`
	Leave      *Leave      `gorm:"foreignKey:LeaveID;comment:假別"`
	Effective  driver.Date `gorm:"type:date;comment:生效時間(yyyy-mm-dd)"`
	Expired    driver.Date `gorm:"type:date;comment:過期時間(yyyy-mm-dd) 空值為下週期自動過期"`
	Available  float64     `gorm:"comment:可用天數"`
	Used       float64     `gorm:"comment:已使用天數"`
	Signing    float64     `gorm:"comment:簽核中天數"`
	ThisTime   float64     `gorm:"comment:這次週期加了幾天"`
	Extra      float64     `gorm:"comment:有幾天是多出來的"`
	PreviousID uint        `gorm:"default:0;not null;comment:上一個遞延ID"`
	NextID     uint        `gorm:"default:0;not null;comment:下一個遞延ID"`
}

func (*LeaveDefer) TableName() string {
	return "leave_defer"
}
