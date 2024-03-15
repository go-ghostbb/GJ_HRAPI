package types

import "gorm.io/gorm"

type PositionRank struct {
	gorm.Model
	Code   string `gorm:"size:100;unique;not null;comment:職級代號" json:"code"`
	Name   string `gorm:"size:100;not null;comment:職級名稱" json:"name"`
	Remark string `gorm:"comment:備註" json:"remark"`

	Grade []*PositionGrade `gorm:"foreignKey:RankID;comment:職等" json:"grade"`
}

func (p *PositionRank) TableName() string {
	return "position_rank"
}
