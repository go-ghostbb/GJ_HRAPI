package types

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Code   string `gorm:"size:100;unique;not null;comment:角色代號"`
	Name   string `gorm:"size:100;not null;comment:工廠名稱"`
	Status bool   `gorm:"default:true;not null;comment:狀態"`
	Remark string `gorm:"comment:備註"`

	Employees   []*Employee   `gorm:"many2many:employee_role;"`
	Permissions []*Permission `gorm:"many2many:role_permission;"`
	Menus       []*Menu       `gorm:"many2many:role_menu;"`
}

func (r *Role) TableName() string {
	return "role"
}
