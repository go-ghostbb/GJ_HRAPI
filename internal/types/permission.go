package types

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Perm   string `gorm:"size:100;not null;unique;comment:權限標示"`
	Name   string `gorm:"size:100;not null;comment:權限名稱"`
	Status bool   `gorm:"default:true;comment:備註"`
	Remark string `gorm:"comment:備註"`

	Roles []*Role `gorm:"many2many:role_permission;"`
}

func (p *Permission) TableName() string {
	return "permission"
}
