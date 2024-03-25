package types

import (
	"gorm.io/gorm"
	"time"
)

type CalcSalary struct {
	gorm.Model
	Start  time.Time `gorm:"comment:開始時間" json:"start"`
	End    time.Time `gorm:"comment:結束時間" json:"end"`
	Salary float32   `gorm:"comment:本薪" json:"salary"`
	Stage  uint      `gorm:"type:tinyint;;comment:當前階段" json:"stage"`

	SalaryAdd    []*CalcSalaryAdd    `gorm:"foreignKey:CalcSalaryID;comment:加項" json:"salaryAdd"`
	SalaryReduce []*CalcSalaryReduce `gorm:"foreignKey:CalcSalaryID;comment:減項" json:"salaryReduce"`
}

func (c *CalcSalary) TableName() string {
	return "calc_salary"
}
