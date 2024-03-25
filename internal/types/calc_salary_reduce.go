package types

import "gorm.io/gorm"

type CalcSalaryReduce struct {
	gorm.Model
	CalcSalaryID uint        `gorm:"not null;comment:calc_salary id" json:"calcSalaryId"`
	CalcSalary   *CalcSalary `gorm:"foreignKey:CalcSalaryID" json:"calcSalary"`

	Name   string  `gorm:"comment:加項名稱" json:"name"`
	Amount float32 `gorm:"comment:金額" json:"amount"`
}

func (c *CalcSalaryReduce) TableName() string {
	return "calc_salary_reduce"
}
