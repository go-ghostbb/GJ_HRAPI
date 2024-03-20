package types

import (
	"gorm.io/gorm"
	"hrapi/internal/types/enum"
)

type OvertimeSignOffSetting struct {
	gorm.Model
	// 部門
	DepartmentID uint        `gorm:"not null;comment:部門ID" json:"departmentId"`
	Department   *Department `gorm:"foreignKey:DepartmentID;comment:部門" json:"department"`

	// 休假日
	VacationID uint      `gorm:"not null;comment:休假日ID, 0為工作日" json:"vacationId"`
	Vacation   *Vacation `gorm:"foreignKey:VacationID;comment:休假日" json:"vacation"`

	Level              uint            `gorm:"type:tinyint;not null;comment:關卡" json:"level"`
	GteHour            float32         `gorm:"column:gte_hour;not null;comment:大於或等於時數" json:"gteHour"`
	SignType           enum.SignType   `gorm:"size:100;comment:特定員工, 部門主管" json:"signType"`
	SpecificEmployeeID uint            `gorm:"comment:特定人員ID" json:"specificEmployeeId"`
	SpecificEmployee   *Employee       `gorm:"foreignKey:SpecificEmployeeID;comment:假別" json:"specificEmployee"`
	Notify             enum.SignNotify `gorm:"size:100;comment:簽核加通知, 僅通知" json:"notify"`
	Remark             string          `gorm:"comment:備註" json:"remark"`
}

func (o *OvertimeSignOffSetting) TableName() string {
	return "overtime_sign_off_setting"
}
