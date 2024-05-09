package types

type EmployeeRole struct {
	RoleID     uint `json:"roleId"`
	EmployeeID uint `json:"employeeId"`
}

func (e *EmployeeRole) TableName() string {
	return "employee_role"
}
