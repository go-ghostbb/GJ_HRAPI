package types

import (
	"gorm.io/gorm"
	"hrapi/internal/types/enum"
	"time"
)

type VacationSchedule struct {
	gorm.Model
	ScheduleDate time.Time `gorm:"type:date;comment:到職日期" json:"scheduleDate"`
	GeneralKey   string    `gorm:"size:129;not null;comment:批量加入的key" json:"generalKey"`
	Remark       string    `gorm:"comment:備註" json:"remark"`

	// Vacation
	VacationID uint      `gorm:"not null;default:0;comment:休假日ID" json:"vacationId"`
	Vacation   *Vacation `gorm:"foreignKey:VacationID;comment:休假日" json:"vacation"`

	// Config
	*VacationScheduleConfig `gorm:"embedded;comment:附加屬性" json:",inline"`
}

func (v *VacationSchedule) TableName() string {
	return "vacation_schedule"
}

type VacationScheduleConfig struct {
	StartDate time.Time                   `gorm:"type:date;comment:開始" json:"startDate"`
	EndDate   time.Time                   `gorm:"type:date;comment:結束" json:"endDate"`
	Repeat    enum.VacationScheduleRepeat `gorm:"size:20;comment:重複類型" json:"repeat"`
	EndRepeat time.Time                   `gorm:"type:date;comment:結束重複日期" json:"endRepeat"`
}
