package types

import "gorm.io/gorm"

type ConfigMap struct {
	gorm.Model
	Key    string `gorm:"comment:key" json:"key"`
	Value  string `gorm:"comment:value" json:"value"`
	Remark string `gorm:"comment:備註" json:"remark"`
}

func (c *ConfigMap) TableName() string {
	return "config_map"
}
