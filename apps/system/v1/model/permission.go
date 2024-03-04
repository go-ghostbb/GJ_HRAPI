package model

import (
	"hrapi/internal/types"
	"hrapi/internal/utils/response/page"
)

type GetByKeywordPermReq struct {
	Keyword string `json:"keyword" form:"keyword"`
	Status  string `json:"status" form:"status"`
}

type GetByKeywordPermRes struct {
	page.Model[types.Permission]
}

type GetByIDPermReq struct {
	ID uint `form:"id"`
}

type GetByIDPermRes struct {
	*types.Permission
}

type PostPermReq struct {
	*types.Permission
}

type PostPermRes struct{}

type PutPermReq struct {
	ID uint
	*types.Permission
}

type PutPermRes struct{}

type DeletePermReq struct {
	ID uint `json:"id" form:"id"`
}

type DeletePermRes struct{}

type SetStatusPermReq struct {
	ID     uint `json:"id" form:"id"`
	Status bool `json:"status" form:"status"`
}

type SetStatusPermRes struct{}
