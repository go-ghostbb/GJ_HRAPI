package types

import (
	"gorm.io/gorm"
	"hrapi/internal/types/enum"
)

type LeaveSignOffSetting struct {
	gorm.Model
	// 部門
	DepartmentID uint        `gorm:"not null;comment:部門ID" json:"departmentId"`
	Department   *Department `gorm:"foreignKey:DepartmentID;comment:部門" json:"department"`

	// 假別
	LeaveID uint   `gorm:"not null;comment:假別ID" json:"leaveId"`
	Leave   *Leave `gorm:"foreignKey:LeaveID;comment:假別" json:"leave"`

	Level              uint            `gorm:"type:tinyint;not null;comment:關卡" json:"level"`
	GteDay             uint            `gorm:"column:gte_day;type:tinyint;not null;comment:大於或等於天數" json:"gteDay"`
	SignType           enum.SignType   `gorm:"size:100;comment:特定員工, 部門主管" json:"signType"`
	SpecificEmployeeID uint            `gorm:"comment:特定人員ID" json:"specificEmployeeId"`
	SpecificEmployee   *Employee       `gorm:"foreignKey:SpecificEmployeeID;comment:假別" json:"specificEmployee"`
	Notify             enum.SignNotify `gorm:"size:100;comment:簽核加通知, 僅通知" json:"notify"`
	Remark             string          `gorm:"comment:備註" json:"remark"`
}

func (l *LeaveSignOffSetting) TableName() string {
	return "leave_sign_off_setting"
}
