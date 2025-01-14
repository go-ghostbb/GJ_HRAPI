package types

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"hrapi/internal/utils/driver"
	"math"
)

type LeaveCorrect struct {
	gorm.Model
	EmployeeID uint        `gorm:"default:0;not null;comment:員工ID" json:"employeeId"`
	Employee   *Employee   `gorm:"foreignKey:EmployeeID;comment:員工" json:"employee"`
	LeaveID    uint        `gorm:"default:0;not null;comment:假別ID" json:"leaveId"`
	Leave      *Leave      `gorm:"foreignKey:LeaveID;comment:假別" json:"leave"`
	Effective  driver.Date `gorm:"type:date;comment:生效時間(yyyy-mm-dd)" json:"effective"`
	Expired    driver.Date `gorm:"type:date;comment:過期時間(yyyy-mm-dd)" json:"expired"`
	Available  float64     `gorm:"comment:可用天數" json:"available"`
	Used       float64     `gorm:"comment:已使用天數" json:"used"`
	Signing    float64     `gorm:"comment:簽核中天數" json:"signing"`
	IsDefer    bool        `gorm:"default:false;comment:是否是遞延" json:"isDefer"`
}

func (LeaveCorrect) TableName() string {
	return "leave_correct"
}

// RoundA 取到整數
// Available = 0.5    => 保留
// Available > 0.5    => 進位
// Available < 0.5    => 捨去
func (l *LeaveCorrect) RoundA() {
	// 提取整數
	intPart := math.Floor(l.Available)
	// 提取小數，使用十進制去做運算，避免二進制進行小數點加減法產生的誤差
	decimalPart, _ := decimal.NewFromFloat(l.Available).Sub(decimal.NewFromFloat(intPart)).Float64()

	switch {
	case decimalPart > 0.5:
		l.Available = intPart + 1
	case decimalPart < 0.5:
		l.Available = intPart
	case decimalPart == 0.5:
		l.Available = intPart + l.Available
	}
}
