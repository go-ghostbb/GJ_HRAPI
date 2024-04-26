package types

import (
	"gorm.io/gorm"
	"hrapi/internal/types/enum"
	"hrapi/internal/utils/driver"
)

type CheckInStatus struct {
	gorm.Model
	EmployeeID              uint                     `gorm:"comment:員工ID" json:"employeeId"`
	Employee                *Employee                `gorm:"foreignKey:EmployeeID;comment:員工" json:"employee"`
	WorkShiftID             uint                     `gorm:"comment:班別ID" json:"workShiftId"`
	WorkShift               *WorkShift               `gorm:"foreignKey:WorkShiftID;comment:班別" json:"workShift"`
	Date                    driver.Date              `gorm:"type:date;comment:上班日期" json:"date"`
	WorkCheckIn             driver.Datetime          `gorm:"type:datetime;comment:上班打卡時間" json:"workCheckIn"`
	WorkAttendStatus        enum.WorkAttendStatus    `gorm:"comment:上班出勤狀態" json:"workAttendStatus"`
	WorkAttendProcStatus    enum.CheckInProcStatus   `gorm:"comment:上班處理狀態" json:"workAttendProcStatus"`
	OffWorkCheckIn          driver.Datetime          `gorm:"type:datetime;comment:下班打卡時間" json:"offWorkCheckIn"`
	OffWorkAttendStatus     enum.OffWorkAttendStatus `gorm:"comment:下班出勤狀態" json:"offWorkAttendStatus"`
	OffWorkAttendProcStatus enum.CheckInProcStatus   `gorm:"comment:下班處理狀態" json:"offWorkAttendProcStatus"`
	AbsenceHours            float32                  `gorm:"comment:缺勤時數" json:"absenceHours"`
	LeaveHours              float32                  `gorm:"comment:請假時數" json:"leaveHours"`
	SignLeaveHours          float32                  `gorm:"comment:簽核中的請假時數" json:"signLeaveHours"`
	OvertimeHours           float32                  `gorm:"comment:加班時數" json:"overtimeHours"`
	SignOvertimeHours       float32                  `gorm:"comment:簽核中的加班時數" json:"signOvertimeHours"`
}

func (c *CheckInStatus) TableName() string {
	return "check_in_status"
}
