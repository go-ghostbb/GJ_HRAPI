package types

import (
	"gorm.io/gorm"
	"hrapi/internal/types/enum"
)

type CalcSalaryAdd struct {
	gorm.Model
	CalcSalaryEmployeeID uint                `gorm:"not null;comment:calc_salary_employee id" json:"calcSalaryEmployeeId"`
	CalcSalaryEmployee   *CalcSalaryEmployee `gorm:"foreignKey:CalcSalaryEmployeeID" json:"calcSalary"`

	SalaryAddItemID uint           `gorm:"comment:加項ID" json:"salaryAddItemId"`
	SalaryAddItem   *SalaryAddItem `gorm:"foreignKey:SalaryAddItemID" json:"salaryAddItem"`

	Amount float32 `gorm:"comment:金額" json:"amount"`

	SalaryType enum.SalaryType `gorm:"size:20;not null;comment:固定薪資或非固定薪資" json:"salaryType"`
	IncomeTax  bool            `gorm:"not null;comment:所得稅" json:"incomeTax"`
	Benefits   bool            `gorm:"not null;comment:職工福利金" json:"benefits"`
	Premiums   bool            `gorm:"not null;comment:補充保費" json:"premiums"`
}

func (c *CalcSalaryAdd) TableName() string {
	return "calc_salary_add"
}
