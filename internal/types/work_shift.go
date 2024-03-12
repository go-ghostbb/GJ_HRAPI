package types

import (
	"gorm.io/gorm"
	"time"
)

type WorkShift struct {
	gorm.Model
	Code       string    `gorm:"size:100;unique;not null;comment:班別代號" json:"code"`
	Name       string    `gorm:"size:100;not null;comment:班別名稱" json:"name"`
	Status     bool      `gorm:"default:true;not null;comment:狀態" json:"status"`
	Remark     string    `gorm:"comment:備註" json:"remark"`
	WorkStart  time.Time `gorm:"type:time(0);not null;comment:上班時間" json:"workStart"`
	WorkEnd    time.Time `gorm:"type:time(0);not null;comment:下班時間" json:"workEnd"`
	RestStart  time.Time `gorm:"type:time(0);not null;comment:休息開始時間" json:"restStart"`
	RestEnd    time.Time `gorm:"type:time(0);not null;comment:休息結束時間" json:"restEnd"`
	TotalHours float64   `gorm:"comment:上班總時數" json:"totalHours"`
	Color      string    `gorm:"size:10;comment:顏色" json:"color"`
}

func (w *WorkShift) TableName() string {
	return "work_shift"
}
