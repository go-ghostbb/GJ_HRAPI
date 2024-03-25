package model

import (
	"hrapi/internal/types"
	"hrapi/internal/utils/response/page"
)

type GetByKeywordSalaryAddItemReq struct {
	Keyword string `form:"keyword"`
}

type GetByKeywordSalaryAddItemRes struct {
	page.Model[types.SalaryAddItem]
}

type PostSalaryAddItemReq struct {
	*types.SalaryAddItem
}

type PostSalaryAddItemRes struct{}

type PutSalaryAddItemReq struct {
	ID uint
	*types.SalaryAddItem
}

type PutSalaryAddItemRes struct{}

type DeleteSalaryAddItemReq struct {
	ID uint
}

type DeleteSalaryAddItemRes struct{}

type GetByKeywordSalaryReduceItemReq struct {
	Keyword string `form:"keyword"`
}

type GetByKeywordSalaryReduceItemRes struct {
	page.Model[types.SalaryReduceItem]
}

type PostSalaryReduceItemReq struct {
	*types.SalaryReduceItem
}

type PostSalaryReduceItemRes struct{}

type PutSalaryReduceItemReq struct {
	ID uint
	*types.SalaryReduceItem
}

type PutSalaryReduceItemRes struct{}

type DeleteSalaryReduceItemReq struct {
	ID uint
}

type DeleteSalaryReduceItemRes struct{}
