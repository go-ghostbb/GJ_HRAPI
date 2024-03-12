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

type WorkShiftApi struct{}

func (w *WorkShiftApi) Init(group *gin.RouterGroup) {
	// 需要有後台權限 (middleware.Software())
	v1 := group.Group("workShift").Use(middleware.Auth(), middleware.Software())
	v1.GET("", middleware.Paginator(), w.getByKeyword)
	v1.GET(":id", w.getByID)
	v1.POST("", w.insert)
	v1.PUT(":id", w.update)
	v1.DELETE(":id", w.delete)
	v1.PATCH(":id/status", w.setStatus)
	// schedule
	v1.GET("schedule", w.getScheduleByDate)
	v1.PUT("schedule/:employeeID/batch", w.updateWorkScheduleBatch)
}

// 根據keyword, status獲取對應班別
//
//	route => /api/v1/workShift
func (w *WorkShiftApi) getByKeyword(c *gin.Context) {
	var (
		ctx  = gbhttp.Ctx(c)
		in   model.GetByKeywordWorkShiftReq
		out  model.GetByKeywordWorkShiftRes
		page = gbhttp.Get(c, "paginator").Interface().(*paginator.Pagination)
		err  error
	)

	if err = gbhttp.ParseQuery(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidQuery, err.Error())
		return
	}

	out, err = service.WorkShift(ctx, page).GetByKeyword(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 根據ID獲取對應班別
//
//	route => /api/v1/workShift/:id
func (w *WorkShiftApi) getByID(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.GetByIDWorkShiftReq
		out model.GetByIDWorkShiftRes
		err error
	)
	in.ID = gbconv.Uint(c.Param("id"))

	out, err = service.WorkShift(ctx).GetByID(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 新增班別
//
//	route => /api/v1/workShift
func (w *WorkShiftApi) insert(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.PostWorkShiftReq
		err error
	)
	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.WorkShift(ctx).Insert(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 更新班別
//
//	route => /api/v1/workShift/:id
func (w *WorkShiftApi) update(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.PutWorkShiftReq
		err error
	)
	in.ID = gbconv.Uint(c.Param("id"))

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.WorkShift(ctx).Update(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 刪除班別
//
//	route => /api/v1/workShift/:id
func (w *WorkShiftApi) delete(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.DeleteWorkShiftReq
		err error
	)
	in.ID = gbconv.Uint(c.Param("id"))

	err = service.WorkShift(ctx).Delete(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 設置狀態
//
//	route => PATCH /api/v1/workShift/:id
func (w *WorkShiftApi) setStatus(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.SetStatusWorkShiftReq
		err error
	)
	in.ID = gbconv.Uint(c.Param("id"))
	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.WorkShift(ctx).SetStatus(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 設定schedule
//
//	route => GET /api/v1/workShift/schedule
func (w *WorkShiftApi) getScheduleByDate(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.GetByDateWorkScheduleReq
		out []*model.GetByDateWorkScheduleRes
		err error
	)

	if err = gbhttp.ParseQuery(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidQuery, err.Error())
		return
	}

	out, err = service.WorkShift(ctx).GetScheduleByDate(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 更新schedule
//
//	route => PUT /api/v1/workShift/schedule/:employeeID/batch
func (w *WorkShiftApi) updateWorkScheduleBatch(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.PutBatchWorkScheduleReq
		err error
	)
	in.EmployeeID = gbconv.Uint(c.Param("employeeID"))

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.WorkShift(ctx).UpdateWorkScheduleBatch(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}
