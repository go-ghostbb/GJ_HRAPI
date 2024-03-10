package types

import "gorm.io/gorm"

type VacationGroup struct {
	gorm.Model
	VacationID uint      `gorm:"not null;comment:休假日ID" json:"vacationId"`
	Vacation   *Vacation `gorm:"foreignKey:VacationID;comment:假別" json:"vacation"`
	Name       string    `gorm:"comment:群組名稱" json:"name"`

	Employee                  []*Employee                  `gorm:"many2many:vacation_group_employee" json:"employee"`
	VacationGroupOvertimeRate []*VacationGroupOvertimeRate `gorm:"foreignKey:VacationGroupID;comment:群組" json:"vacationGroupOvertimeRate"`
}

func (v *VacationGroup) TableName() string {
	return "vacation_group"
}
