package types

import (
	"gorm.io/gorm"
	"hrapi/internal/types/enum"
)

type LeaveGroupCondition struct {
	gorm.Model
	LeaveGroupID uint        `gorm:"not null;comment:群組ID" json:"leaveGroupId"`
	LeaveGroup   *LeaveGroup `gorm:"foreignKey:LeaveGroupID;comment:群組" json:"leaveGroup"`
	// Condition
	// 假設k為年資(單位:月)
	// 1: k >= m (年資滿m個月)
	// 2: k >= m*12 (年資滿m年)
	ConditionType enum.LeaveGroupCondition `gorm:"comment:條件 1:年資滿<#1>個月 2:年資滿<#1>年" json:"conditionType"`
	ConditionNum  uint                     `gorm:"comment:<#1>" json:"conditionNum"`
	Result        uint                     `gorm:"comment:給予天數" json:"result"`
	Level         uint                     `gorm:"type:tinyint;comment:關卡" json:"level"`
}

func (l *LeaveGroupCondition) TableName() string {
	return "leave_group_condition"
}
