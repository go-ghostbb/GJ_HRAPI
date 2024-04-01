package types

type SalaryAddItemEmployee struct {
	SalaryAddItemID uint    `gorm:"primarykey" json:"salaryAddItemId"`
	EmployeeID      uint    `gorm:"primarykey" json:"employeeId"`
	UseCustom       bool    `json:"useCustom"`
	CustomAmount    float32 `json:"customAmount"`
}

func (s *SalaryAddItemEmployee) TableName() string {
	return "salary_add_item_employee"
}
