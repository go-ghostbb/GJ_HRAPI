package types

import (
	"gorm.io/gorm"
	"hrapi/internal/types/enum"
)

type SalaryReduceItem struct {
	gorm.Model
	Name      string `gorm:"size:100;not null;comment:名稱" json:"name"`
	IncomeTax bool   `gorm:"not null;comment:所得稅" json:"incomeTax"`

	// --------金額設定--------
	Amount    float32             `gorm:"not null;comment:金額" json:"amount"`
	CalcType  enum.SalaryCalcType `gorm:"not null;comment:計算類型" json:"calcType"`
	MonthCalc bool                `gorm:"not null;comment:破月計算, 未滿足月時按比例計算, 預設金額用" json:"monthCalc"`
	Unit      enum.SalaryCalcUnit `gorm:"comment:單位：天, 小時" json:"unit"`
	Operator  enum.Operator       `gorm:"comment:運算子, 條件用(>, <, <=...)" json:"operator"`
	Argument  float32             `gorm:"comment:條件參數" json:"argument"`
}

func (s *SalaryReduceItem) TableName() string {
	return "salary_reduce_item"
}
