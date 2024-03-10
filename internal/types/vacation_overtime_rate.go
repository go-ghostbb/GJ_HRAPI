package types

import "gorm.io/gorm"

type VacationGroupOvertimeRate struct {
	gorm.Model
	VacationGroupID uint        `gorm:"not null;comment:群組ID" json:"vacationGroupId"`
	VacationGroup   *LeaveGroup `gorm:"foreignKey:VacationGroupID;comment:群組" json:"vacationGroup"`

	Hours    uint    `gorm:"type:tinyint;comment:第幾個小時" json:"hours"`
	Multiply float32 `gorm:"comment:倍率" json:"multiply"`
	Level    uint    `gorm:"type:tinyint;comment:關卡" json:"level"`
	IsFill   bool    `gorm:"default:false;comment:是否填滿" json:"isFill"`
	Fill     uint    `gorm:"type:tinyint;comment:自動填滿幾小時 注意：不能為0和大於等於下一個區間的hours" json:"fill"`
}

func (v *VacationGroupOvertimeRate) TableName() string {
	return "vacation_group_overtime_rate"
}
