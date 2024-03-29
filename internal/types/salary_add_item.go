package types

import (
	"gorm.io/gorm"
	"hrapi/internal/types/enum"
)

type SalaryAddItem struct {
	gorm.Model
	Name       string          `gorm:"size:100;not null;comment:名稱" json:"name"`
	SalaryType enum.SalaryType `gorm:"size:20;not null;comment:固定薪資或非固定薪資" json:"salaryType"`
	IncomeTax  bool            `gorm:"not null;comment:所得稅" json:"incomeTax"`
	Benefits   bool            `gorm:"not null;comment:職工福利金" json:"benefits"`
	Premiums   bool            `gorm:"not null;comment:補充保費" json:"premiums"`

	// --------金額設定--------
	Amount    float32             `gorm:"not null;comment:金額" json:"amount"`
	CalcType  enum.SalaryCalcType `gorm:"not null;comment:計算類型" json:"calcType"`
	MonthCalc bool                `gorm:"not null;comment:破月計算, 未滿足月時按比例計算, 預設金額用" json:"monthCalc"`
	Unit      enum.SalaryCalcUnit `gorm:"comment:單位：天, 小時" json:"unit"`
	Operator  enum.Operator       `gorm:"comment:運算子, 條件用(>, <, <=...)" json:"operator"`
	Argument  float32             `gorm:"comment:條件參數" json:"argument"`

	Employee []*Employee `gorm:"many2many:salary_add_item_employee" json:"employee"`
}

func (s *SalaryAddItem) TableName() string {
	return "salary_add_item"
}
