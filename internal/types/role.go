package types

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Code   string `gorm:"size:100;unique;not null;comment:角色代號" json:"code"`
	Name   string `gorm:"size:100;not null;comment:角色名稱" json:"name"`
	Status bool   `gorm:"default:true;not null;comment:狀態" json:"status"`
	Remark string `gorm:"comment:備註" json:"remark"`

	Employees   []*Employee   `gorm:"many2many:employee_role;" json:"employees"`
	Permissions []*Permission `gorm:"many2many:role_permission;" json:"permissions"`
	Menus       []*Menu       `gorm:"many2many:role_menu;" json:"menus"`
}

func (r *Role) TableName() string {
	return "role"
}
