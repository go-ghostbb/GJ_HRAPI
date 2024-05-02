package types

import (
	"gorm.io/gorm"
	"hrapi/internal/types/enum"
	"hrapi/internal/utils/driver"
)

type VacationSchedule struct {
	gorm.Model
	ScheduleDate driver.Date `gorm:"type:date;comment:到職日期" json:"scheduleDate"`
	GeneralKey   string      `gorm:"size:129;not null;comment:批量加入的key" json:"generalKey"`
	Remark       string      `gorm:"comment:備註" json:"remark"`

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
	StartDate driver.Date                 `gorm:"type:date;comment:開始" json:"startDate"`
	EndDate   driver.Date                 `gorm:"type:date;comment:結束" json:"endDate"`
	Repeat    enum.VacationScheduleRepeat `gorm:"size:20;comment:重複類型" json:"repeat"`
	EndRepeat driver.Date                 `gorm:"type:date;comment:結束重複日期" json:"endRepeat"`
}
