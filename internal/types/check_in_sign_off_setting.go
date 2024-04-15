package types

import (
	"gorm.io/gorm"
	"hrapi/internal/types/enum"
)

type CheckInSignOffSetting struct {
	gorm.Model

	Level              uint            `gorm:"type:tinyint;not null;comment:關卡" json:"level"`
	SignType           enum.SignType   `gorm:"size:100;comment:特定員工, 部門主管" json:"signType"`
	SpecificEmployeeID uint            `gorm:"comment:特定人員ID" json:"specificEmployeeId"`
	SpecificEmployee   *Employee       `gorm:"foreignKey:SpecificEmployeeID;comment:假別" json:"specificEmployee"`
	Notify             enum.SignNotify `gorm:"size:100;comment:簽核加通知, 僅通知" json:"notify"`
	Remark             string          `gorm:"comment:備註" json:"remark"`
}

func (c *CheckInSignOffSetting) TableName() string {
	return "check_in_sign_off_setting"
}
