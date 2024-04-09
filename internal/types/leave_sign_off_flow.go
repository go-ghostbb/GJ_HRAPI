package types

import (
	"gorm.io/gorm"
	"hrapi/internal/types/enum"
	"time"
)

type LeaveSignOffFlow struct {
	gorm.Model
	LeaveRequestFormID uint              `gorm:"not null;comment:請假單ID"`
	LeaveRequestForm   *LeaveRequestForm `gorm:"foreignKey:LeaveRequestFormID;comment:請假單"`
	SignOffEmployeeID  uint              `gorm:"not null;comment:簽核人員ID"`
	SignOffEmployee    *Employee         `gorm:"foreignKey:SignOffEmployeeID;comment:簽核人員"`
	Level              uint              `gorm:"comment:關卡"`
	SignType           enum.SignType     `gorm:"comment:1:部門主管 2:特定人員 3:代理人 4:起單人"`
	Notify             enum.SignNotify   `gorm:"size:100;comment:1:簽核通知 2:僅通知"`
	Remark             string            `gorm:"comment:備註'"`
	Comment            string            `gorm:"comment:簽核意見"`
	Status             enum.SignStatus   `gorm:"type:tinyint;comment:1:核准 2:駁回 3:僅通知，通知成功 4:簽核中 5:僅通知，通知失敗"`
	SignDate           time.Time         `gorm:"comment:覆合日期"`
	UUID               string            `gorm:"comment:UUID"`
	Locale             enum.Locale       `gorm:"size:20;comment:語言"`
}

func (l *LeaveSignOffFlow) TableName() string {
	return "leave_sign_off_flow"
}
