package types

import "gorm.io/gorm"

type Vacation struct {
	gorm.Model
	Code   string `gorm:"size:100;unique;not null;comment:休假日類型代號" json:"code"`
	Name   string `gorm:"size:100;not null;comment:休假日類型名稱" json:"name"`
	Status bool   `gorm:"default:true;not null;comment:狀態" json:"status"`
	Remark string `gorm:"comment:備註" json:"remark"`
	Color  string `gorm:"size:10;comment:顏色" json:"color"`
	Weight int    `gorm:"default:1;comment:權重" json:"weight"`

	// schedule
	Schedule []*VacationSchedule `gorm:"foreignKey:VacationID" json:"schedule"`

	// overtime rate
	VacationGroup []*VacationGroup `gorm:"foreignKey:VacationID" json:"overtimeRate"`
}

func (v *Vacation) TableName() string {
	return "vacation"
}
