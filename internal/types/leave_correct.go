package types

import "gorm.io/gorm"

type LeaveCorrect struct {
	gorm.Model
	EmployeeID uint      `gorm:"default:0;not null;comment:員工ID"`
	Employee   *Employee `gorm:"foreignKey:EmployeeID;comment:員工"`
	LeaveID    uint      `gorm:"default:0;not null;comment:假別ID"`
	Leave      *Leave    `gorm:"foreignKey:LeaveID;comment:假別"`
	Available  float64   `gorm:"comment:可用天數"`
	Used       float64   `gorm:"comment:已使用天數"`
	Signing    float64   `gorm:"comment:簽核中天數"`
}

func (LeaveCorrect) TableName() string {
	return "leave_correct"
}
