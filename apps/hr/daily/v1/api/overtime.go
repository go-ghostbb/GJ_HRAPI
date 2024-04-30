package api

import (
	gbhttp "ghostbb.io/gb/net/gb_http"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"github.com/gin-gonic/gin"
	"hrapi/apps/hr/daily/v1/model"
	"hrapi/apps/hr/daily/v1/service"
	"hrapi/internal/middleware"
	"hrapi/internal/utils/paginator"

	. "hrapi/internal/utils/response"
	. "hrapi/internal/utils/response/status"
)

type OvertimeApi struct{}

func (ch *OvertimeApi) Init(group *gin.RouterGroup) {
	// 需要有網頁端權限 (middleware.Web())
	v1 := group.Group("daily/overtime").Use(middleware.Auth(), middleware.Web())
	v1.GET("", middleware.Paginator(), ch.getOvertimeFormByEmployee)
	v1.POST("", ch.saveOvertimeForm)
	v1.DELETE(":id", ch.deleteOvertimeForm)
}

// 獲取請假單列表
//
//	route => GET /api/v1/daily/overtime
func (ch *OvertimeApi) getOvertimeFormByEmployee(c *gin.Context) {
	var (
		ctx  = gbhttp.Ctx(c)
		in   model.GetOvertimeFormByEmployeeReq
		out  model.GetOvertimeFormByEmployeeRes
		page = gbhttp.Get(c, "paginator").Interface().(*paginator.Pagination)
		err  error
	)

	in.EmployeeID = gbhttp.Get(c, "employee.id").Uint()

	if err = gbhttp.ParseQuery(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidQuery, err.Error())
		return
	}

	out, err = service.Overtime(ctx, page).GetOvertimeFormByEmployee(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 建立或更新補打卡單
//
//	route => POST /api/v1/daily/overtime
func (ch *OvertimeApi) saveOvertimeForm(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.SaveOvertimeFormReq
		out model.SaveOvertimeFormRes
		err error
	)

	in.EmployeeID = gbhttp.Get(c, "employee.id").Uint()

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	out, err = service.Overtime(ctx).SaveOvertimeForm(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 刪除補打卡單
//
//	route => DELETE /api/v1/daily/overtime/:id
func (ch *OvertimeApi) deleteOvertimeForm(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.DeleteOvertimeFormReq
		err error
	)

	in.ID = gbconv.Uint(c.Param("id"))

	err = service.Overtime(ctx).DeleteOvertimeForm(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}
