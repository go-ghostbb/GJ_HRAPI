package model

import "hrapi/internal/types"

type GetByKeywordMenuReq struct {
	Keyword string `form:"keyword" json:"keyword"`
	Status  string `form:"status" json:"status"`
	Show    string `form:"show" json:"show"`
}

type GetByKeywordMenuRes struct {
	*types.Menu `json:",inline"`
}

type GetByIDMenuReq struct {
	ID uint `json:"id"`
}

type GetByIDMenuRes struct {
	*types.Menu `json:",inline"`
}

type PostMenuReq struct {
	*types.Menu `json:",inline"`
}

type PostMenuRes struct{}

type PutMenuReq struct {
	ID          uint `form:"id"`
	*types.Menu `json:",inline"`
}

type PutMenuRes struct{}

type DeleteMenuReq struct {
	ID uint `json:"id"`
}

type DeleteMenuRes struct{}

type SetStatusMenuReq struct {
	ID     uint
	Status bool `json:"status" form:"status"`
}

type SetStatusMenuRes struct{}
