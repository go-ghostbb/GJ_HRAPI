package model

import (
	"hrapi/internal/types"
	"hrapi/internal/utils/response/page"
)

type GetByKeywordGradeReq struct {
	RankID  uint   `json:"rankId" binding:"required" form:"rankId"`
	Keyword string `json:"keyword" form:"keyword"`
}

type GetByKeywordGradeRes struct {
	page.Model[types.PositionGrade]
}

type GetByIDGradeReq struct {
	ID uint `json:"id" form:"id"`
}

type GetByIDGradeRes struct {
	*types.PositionGrade
}

type PostGradeReq struct {
	*types.PositionGrade
}

type PostGradeRes struct{}

type PutGradeReq struct {
	ID uint
	*types.PositionGrade
}

type PutGradeRes struct{}

type DeleteGradeReq struct {
	ID uint
}

type DeleteGradeRes struct{}
