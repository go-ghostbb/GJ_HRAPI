package model

import (
	"hrapi/internal/types"
	"hrapi/internal/utils/response/page"
)

type GetByKeywordRankReq struct {
	Keyword string `json:"keyword" form:"keyword"`
}

type GetByKeywordRankRes struct {
	page.Model[types.PositionRank]
}

type GetByIDRankReq struct {
	ID uint `json:"id" form:"id"`
}

type GetByIDRankRes struct {
	*types.PositionRank
}

type PostRankReq struct {
	*types.PositionRank
}

type PostRankRes struct{}

type PutRankReq struct {
	ID uint
	*types.PositionRank
}

type PutRankRes struct{}

type DeleteRankReq struct {
	ID uint
}

type DeleteRankRes struct{}
