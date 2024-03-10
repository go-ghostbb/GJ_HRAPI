package types

type LeaveGroupEmployee struct {
	LeaveGroupID uint `gorm:"column:leave_group_id"`
	EmployeeID   uint `gorm:"column:employee_id"`
}

func (l *LeaveGroupEmployee) TableName() string {
	return "leave_group_employee"
}
