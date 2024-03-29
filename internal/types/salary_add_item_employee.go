package types

type SalaryAddItemEmployee struct {
	SalaryAddItemID uint `json:"salaryAddItemId"`
	EmployeeID      uint `json:"employeeId"`
}

func (s *SalaryAddItemEmployee) TableName() string {
	return "salary_add_item_employee"
}
