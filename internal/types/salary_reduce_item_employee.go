package types

type SalaryReduceItemEmployee struct {
	SalaryReduceItemID uint `json:"salaryReduceItemId"`
	EmployeeID         uint `json:"employeeId"`
}

func (s *SalaryReduceItemEmployee) TableName() string {
	return "salary_reduce_item_employee"
}
