package types

import "gorm.io/gorm"

type LeaveGroup struct {
	gorm.Model
	LeaveID uint   `gorm:"not null;comment:假別ID" json:"leaveId"`
	Leave   *Leave `gorm:"foreignKey:LeaveID;comment:假別" json:"leave"`
	Name    string `gorm:"comment:群組名稱" json:"name"`

	Employee            []*Employee            `gorm:"many2many:leave_group_employee" json:"employee"`
	LeaveGroupCondition []*LeaveGroupCondition `gorm:"foreignKey:LeaveGroupID;comment:群組" json:"leaveGroupCondition"`
}

func (l *LeaveGroup) TableName() string {
	return "leave_group"
}
