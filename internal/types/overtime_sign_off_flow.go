package types

import (
	"gorm.io/gorm"
	"hrapi/internal/types/enum"
	"time"
)

type OvertimeSignOffFlow struct {
	gorm.Model
	OvertimeRequestFormID uint                 `gorm:"not null;comment:加班ID" json:"overtimeRequestFormId"`
	OvertimeRequestForm   *OvertimeRequestForm `gorm:"foreignKey:OvertimeRequestFormID;comment:加班單" json:"overtimeRequestForm"`
	SignOffEmployeeID     uint                 `gorm:"not null;comment:簽核人員ID" json:"signOffEmployeeId"`
	SignOffEmployee       *Employee            `gorm:"foreignKey:SignOffEmployeeID;comment:簽核人員" json:"signOffEmployee"`
	Level                 uint                 `gorm:"comment:關卡" json:"level"`
	SignType              enum.SignType        `gorm:"comment:1:部門主管 2:特定人員 3:代理人 4:起單人" json:"signType"`
	Notify                enum.SignNotify      `gorm:"size:100;comment:1:簽核通知 2:僅通知" json:"notify"`
	Remark                string               `gorm:"comment:備註'" json:"remark"`
	Comment               string               `gorm:"comment:簽核意見" json:"comment"`
	Status                enum.SignStatus      `gorm:"type:tinyint;comment:1:核准 2:駁回 3:僅通知，通知成功 4:簽核中 5:僅通知，通知失敗" json:"status"`
	SignDate              time.Time            `gorm:"comment:覆合日期" json:"signDate"`
	UUID                  string               `gorm:"comment:UUID" json:"uuid"`
	Locale                enum.Locale          `gorm:"size:20;comment:語言" json:"locale"`
}

func (l *OvertimeSignOffFlow) TableName() string {
	return "overtime_sign_off_flow"
}
