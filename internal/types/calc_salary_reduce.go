package types

import "gorm.io/gorm"

type CalcSalaryReduce struct {
	gorm.Model
	CalcSalaryEmployeeID uint                `gorm:"not null;comment:calc_salary_employee id" json:"calcSalaryEmployeeId"`
	CalcSalaryEmployee   *CalcSalaryEmployee `gorm:"foreignKey:CalcSalaryEmployeeID" json:"calcSalary"`

	SalaryReduceItemID uint              `gorm:"comment:加項ID" json:"salaryReduceItemId"`
	SalaryReduceItem   *SalaryReduceItem `gorm:"foreignKey:SalaryReduceItemID" json:"salaryReduceItem"`

	Amount float32 `gorm:"comment:金額" json:"amount"`

	IncomeTax bool `gorm:"not null;comment:所得稅" json:"incomeTax"`
}

func (c *CalcSalaryReduce) TableName() string {
	return "calc_salary_reduce"
}
