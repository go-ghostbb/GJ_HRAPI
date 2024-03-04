package types

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Perm   string `gorm:"size:100;not null;unique;comment:權限標示" json:"perm"`
	Name   string `gorm:"size:100;not null;comment:權限名稱" json:"name"`
	Status bool   `gorm:"default:true;comment:備註" json:"status"`
	Remark string `gorm:"comment:備註" json:"remark"`

	Roles []*Role `gorm:"many2many:role_permission;" json:"roles"`
}

func (p *Permission) TableName() string {
	return "permission"
}
