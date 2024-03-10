package types

import (
	"gorm.io/gorm"
	"time"
)

type VacationSchedule struct {
	gorm.Model
	ScheduleDate time.Time `gorm:"type:date;comment:到職日期" json:"scheduleDate"`

	// Employee
	EmployeeID uint      `gorm:"not null;default:0;comment:員工ID" json:"employeeId"`
	Employee   *Employee `gorm:"foreignKey:EmployeeID;comment:員工" json:"employee"`

	// Vacation
	VacationID uint      `gorm:"not null;default:0;comment:休假日ID" json:"vacationId"`
	Vacation   *Vacation `gorm:"foreignKey:VacationID;comment:休假日" json:"vacation"`
}

func (v *VacationSchedule) TableName() string {
	return "vacation_schedule"
}
