package types

import "gorm.io/gorm"

type CalcSalaryEmployee struct {
	gorm.Model
	EmployeeID uint      `gorm:"not null;comment:員工ID" json:"employeeId"`
	Employee   *Employee `gorm:"foreignKey:EmployeeID" json:"employee"`

	CalcSalaryID uint        `gorm:"not null;comment:calc_salary id" json:"calcSalaryId"`
	CalcSalary   *CalcSalary `gorm:"foreignKey:CalcSalaryID" json:"calcSalary"`

	Salary       float32 `gorm:"comment:本薪" json:"salary"`
	HourlySalary float32 `gorm:"comment:時薪" json:"hourlySalary"`

	SalaryAdd    []*CalcSalaryAdd    `gorm:"foreignKey:CalcSalaryEmployeeID;comment:加項" json:"salaryAdd"`
	SalaryReduce []*CalcSalaryReduce `gorm:"foreignKey:CalcSalaryEmployeeID;comment:減項" json:"salaryReduce"`
}

func (c *CalcSalaryEmployee) TableName() string {
	return "calc_salary_employee"
}
