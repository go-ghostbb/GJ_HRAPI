package api

import (
	gbhttp "ghostbb.io/gb/net/gb_http"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"github.com/gin-gonic/gin"
	"hrapi/apps/hr/setting/v1/model"
	"hrapi/apps/hr/setting/v1/service"
	"hrapi/internal/middleware"
	"hrapi/internal/utils/paginator"

	. "hrapi/internal/utils/response"
	. "hrapi/internal/utils/response/status"
)

type SalaryItemApi struct{}

func (s *SalaryItemApi) Init(group *gin.RouterGroup) {
	// 需要有後台權限 (middleware.Software())
	v1 := group.Group("salary").Use(middleware.Auth(), middleware.Software())
	v1.GET("add", middleware.Paginator(), s.getByKeywordAdd)
	v1.POST("add", s.insertAdd)
	v1.PUT("add/:id", s.updateAdd)
	v1.DELETE("add/:id", s.deleteAdd)

	v1.GET("reduce", middleware.Paginator(), s.getByKeywordReduce)
	v1.POST("reduce", s.insertReduce)
	v1.PUT("reduce/:id", s.updateReduce)
	v1.DELETE("reduce/:id", s.deleteReduce)
}

// 根據keyword獲取薪資加項設定
//
//	route => GET /api/v1/salary/add
func (s *SalaryItemApi) getByKeywordAdd(c *gin.Context) {
	var (
		ctx  = gbhttp.Ctx(c)
		in   model.GetByKeywordSalaryAddItemReq
		out  model.GetByKeywordSalaryAddItemRes
		page = gbhttp.Get(c, "paginator").Interface().(*paginator.Pagination)
		err  error
	)
	if err = gbhttp.ParseQuery(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidQuery, err.Error())
		return
	}

	out, err = service.SalaryItem(ctx, page).GetByKeywordAdd(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 新增薪資加項
//
//	route => POST /api/v1/salary/add
func (s *SalaryItemApi) insertAdd(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.PostSalaryAddItemReq
		err error
	)
	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.SalaryItem(ctx).InsertAdd(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 更新薪資加項
//
//	route => PUT /api/v1/salary/add/:id
func (s *SalaryItemApi) updateAdd(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.PutSalaryAddItemReq
		err error
	)
	in.ID = gbconv.Uint(c.Param("id"))

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.SalaryItem(ctx).UpdateAdd(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 刪除薪資加項
//
//	route => DELETE /api/v1/salary/add/:id
func (s *SalaryItemApi) deleteAdd(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.DeleteSalaryAddItemReq
		err error
	)

	in.ID = gbconv.Uint(c.Param("id"))

	err = service.SalaryItem(ctx).DeleteAdd(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 根據keyword獲取薪資減項設定
//
//	route => GET /api/v1/salary/reduce
func (s *SalaryItemApi) getByKeywordReduce(c *gin.Context) {
	var (
		ctx  = gbhttp.Ctx(c)
		in   model.GetByKeywordSalaryReduceItemReq
		out  model.GetByKeywordSalaryReduceItemRes
		page = gbhttp.Get(c, "paginator").Interface().(*paginator.Pagination)
		err  error
	)
	if err = gbhttp.ParseQuery(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidQuery, err.Error())
		return
	}

	out, err = service.SalaryItem(ctx, page).GetByKeywordReduce(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 新增薪資減項
//
//	route => POST /api/v1/salary/reduce
func (s *SalaryItemApi) insertReduce(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.PostSalaryReduceItemReq
		err error
	)
	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.SalaryItem(ctx).InsertReduce(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 更新薪資減項
//
//	route => PUT /api/v1/salary/reduce/:id
func (s *SalaryItemApi) updateReduce(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.PutSalaryReduceItemReq
		err error
	)
	in.ID = gbconv.Uint(c.Param("id"))

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.SalaryItem(ctx).UpdateReduce(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 刪除薪資減項
//
//	route => DELETE /api/v1/salary/reduce/:id
func (s *SalaryItemApi) deleteReduce(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.DeleteSalaryReduceItemReq
		err error
	)

	in.ID = gbconv.Uint(c.Param("id"))

	err = service.SalaryItem(ctx).DeleteReduce(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}
