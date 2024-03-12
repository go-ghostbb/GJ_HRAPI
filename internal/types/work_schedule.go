package types

import (
	"gorm.io/gorm"
	"time"
)

type WorkSchedule struct {
	gorm.Model
	ScheduleDate time.Time `gorm:"type:date;comment:到職日期" json:"scheduleDate"`
	// User
	EmployeeID uint      `gorm:"not null;default:0;comment:員工ID" json:"employeeId"`
	Employee   *Employee `gorm:"foreignKey:EmployeeID;comment:員工" json:"employee"`

	// WorkShift
	WorkShiftID uint       `gorm:"not null;default:0;comment:班別ID" json:"workShiftId"`
	WorkShift   *WorkShift `gorm:"foreignKey:WorkShiftID;comment:班別" json:"workShift"`
}

func (w *WorkSchedule) TableName() string {
	return "work_schedule"
}
