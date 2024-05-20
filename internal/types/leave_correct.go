package types

import (
	"gorm.io/gorm"
	"hrapi/internal/utils/driver"
)

type LeaveCorrect struct {
	gorm.Model
	EmployeeID uint        `gorm:"default:0;not null;comment:員工ID" json:"employeeId"`
	Employee   *Employee   `gorm:"foreignKey:EmployeeID;comment:員工" json:"employee"`
	LeaveID    uint        `gorm:"default:0;not null;comment:假別ID" json:"leaveId"`
	Leave      *Leave      `gorm:"foreignKey:LeaveID;comment:假別" json:"leave"`
	Effective  driver.Date `gorm:"type:date;comment:生效時間(yyyy-mm-dd)" json:"effective"`
	Expired    driver.Date `gorm:"type:date;comment:過期時間(yyyy-mm-dd)" json:"expired"`
	Available  float64     `gorm:"comment:可用天數" json:"available"`
	Used       float64     `gorm:"comment:已使用天數" json:"used"`
	Signing    float64     `gorm:"comment:簽核中天數" json:"signing"`
	IsDefer    bool        `gorm:"default:false;comment:是否是遞延" json:"isDefer"`
}

func (LeaveCorrect) TableName() string {
	return "leave_correct"
}
