package types

import "gorm.io/gorm"

type OvertimeRequestFormRate struct {
	gorm.Model
	OvertimeRequestFormID uint                 `gorm:"not null;comment:加班ID" json:"overtimeRequestFormId"`
	OvertimeRequestForm   *OvertimeRequestForm `gorm:"foreignKey:OvertimeRequestFormID;comment:加班單" json:"overtimeRequestForm"`
	Hours                 uint                 `gorm:"type:tinyint;comment:第幾個小時" json:"hours"`
	Multiply              float32              `gorm:"comment:倍率" json:"multiply"`
	Level                 uint                 `gorm:"type:tinyint;comment:關卡" json:"level"`
	IsFill                bool                 `gorm:"default:false;comment:是否填滿" json:"isFill"`
	Fill                  uint                 `gorm:"type:tinyint;comment:自動填滿幾小時 注意：不能為0和大於等於下一個區間的hours" json:"fill"`
}

func (o *OvertimeRequestFormRate) TableName() string {
	return "overtime_request_form_rate"
}
