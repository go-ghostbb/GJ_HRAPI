package types

import (
	"gorm.io/gorm"
	"time"
)

type CalcSalary struct {
	gorm.Model

	FounderEmployeeID uint      `gorm:"comment:創建人ID" json:"founderEmployeeId"`
	Founder           *Employee `gorm:"foreignKey:FounderEmployeeID;comment:創建人" json:"founder"`

	Start time.Time `gorm:"type:date;comment:開始時間" json:"start"`
	End   time.Time `gorm:"type:date;comment:結束時間" json:"end"`
	Stage uint      `gorm:"type:tinyint;;comment:當前階段" json:"stage"`

	CalcSalaryEmployee []*CalcSalaryEmployee `gorm:"foreignKey:CalcSalaryID;comment:員工" json:"calcSalaryEmployee"`
}

func (c *CalcSalary) TableName() string {
	return "calc_salary"
}
