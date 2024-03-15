package types

import "gorm.io/gorm"

type PositionGrade struct {
	gorm.Model
	Code   string `gorm:"size:100;unique;not null;comment:職等代號" json:"code"`
	Name   string `gorm:"size:100;not null;comment:職等名稱" json:"name"`
	Remark string `gorm:"comment:備註" json:"remark"`

	RankID uint          `gorm:"not null;comment:職級ID" json:"rankId"`
	Rank   *PositionRank `gorm:"foreignKey:RankID;comment:職級" json:"rank"`
}

func (p *PositionGrade) TableName() string {
	return "position_grade"
}
