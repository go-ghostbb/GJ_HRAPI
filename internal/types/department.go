package types

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	ParentID uint   `gorm:"default:0;not null;comment:父部門" json:"parentId"`
	Code     string `gorm:"size:100;unique;not null;comment:部門代號" json:"code"`
	Name     string `gorm:"size:100;not null;comment:部門名稱" json:"name"`
	Status   bool   `gorm:"default:true;not null;comment:狀態" json:"status"`
	Remark   string `gorm:"comment:備註" json:"remark"`

	ManagerId uint      `gorm:"default:0;not null;comment:主管ID" json:"managerId"`
	Manager   *Employee `gorm:"foreignKey:ManagerId;comment:主管" json:"manager"`

	Children []*Department `gorm:"-" json:"children"`
}

func (d *Department) TableName() string {
	return "department"
}
