package types

import (
	"gorm.io/gorm"
	"hrapi/internal/utils/driver"
)

type LeaveCorrect struct {
	gorm.Model
	EmployeeID uint        `gorm:"default:0;not null;comment:員工ID"`
	Employee   *Employee   `gorm:"foreignKey:EmployeeID;comment:員工"`
	LeaveID    uint        `gorm:"default:0;not null;comment:假別ID"`
	Leave      *Leave      `gorm:"foreignKey:LeaveID;comment:假別"`
	Effective  driver.Date `gorm:"type:date;comment:生效時間(yyyy-mm-dd)"`
	Expired    driver.Date `gorm:"type:date;comment:過期時間(yyyy-mm-dd)"`
	Available  float64     `gorm:"comment:可用天數"`
	Used       float64     `gorm:"comment:已使用天數"`
	Signing    float64     `gorm:"comment:簽核中天數"`
	IsDefer    bool        `gorm:"default:false;comment:是否是遞延"`
}

func (LeaveCorrect) TableName() string {
	return "leave_correct"
}
