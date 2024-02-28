package types

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	ParentID uint   `gorm:"default:0;not null;comment:父部門"`
	Code     string `gorm:"size:100;unique;not null;comment:部門代號"`
	Name     string `gorm:"size:100;not null;comment:部門名稱"`
	Status   bool   `gorm:"default:true;not null;comment:狀態"`
	Remark   string `gorm:"comment:備註"`

	ManagerId uint      `gorm:"default:0;not null;comment:主管ID"`
	Manager   *Employee `gorm:"foreignKey:ManagerId;comment:主管"`
}
