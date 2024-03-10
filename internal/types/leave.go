package types

import (
	"gorm.io/gorm"
	"hrapi/internal/types/enum"
)

type Leave struct {
	gorm.Model
	Code    string          `gorm:"size:100;unique;not null;comment:代碼" json:"code"`
	Name    string          `gorm:"size:100;not null;comment:請假名稱" json:"name"`
	Status  bool            `gorm:"default:true;not null;comment:狀態" json:"status"`
	Default uint            `gorm:"default:0;comment:預設給假(單位：天)" json:"default"`
	Minimum uint            `gorm:"comment:最小請假時間(單位：分鐘)" json:"minimum"`
	Pay     enum.LeavePay   `gorm:"default:none;not null;comment:全薪, 半薪, 不給薪" json:"pay"`
	Cycle   enum.LeaveCycle `gorm:"not null;comment:預設(以「年」為單位給予), 週年制, 歷年制(曆年制" json:"cycle"`
	Remark  string          `gorm:"comment:備註" json:"remark"`

	// -------遞延相關資料-------
	Deferrable           bool `gorm:"default:false;comment:是否可遞延" json:"deferrable"`
	DeferrableDays       uint `gorm:"comment:可遞延幾天" json:"deferrableDays"`
	CustomizableDuration bool `gorm:"default:false;comment:是否可自訂到期時間" json:"customizableDuration"`
	Duration             uint `gorm:"default:0;comment:遞延持續時間(單位：月)" json:"duration"`
	// ------------------------

	LeaveGroup []*LeaveGroup `gorm:"foreignKey:LeaveID;comment:群組" json:"leaveGroup"`
}

func (l *Leave) TableName() string {
	return "leave"
}
