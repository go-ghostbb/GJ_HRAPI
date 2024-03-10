package types

type VacationGroupEmployee struct {
	VacationGroupID uint `gorm:"column:vacation_group_id"`
	EmployeeID      uint `gorm:"column:employee_id"`
}

func (l *VacationGroupEmployee) TableName() string {
	return "vacation_group_employee"
}
