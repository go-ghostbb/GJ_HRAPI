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

type LeaveApi struct{}

func (l *LeaveApi) Init(group *gin.RouterGroup) {
	// 需要有網頁端權限 (middleware.Web())
	v1 := group.Group("daily/leave").Use(middleware.Auth(), middleware.Web())
	v1.GET("", middleware.Paginator(), l.getLeaveFormByEmployee)
	v1.POST("", l.saveLeaveForm)
	v1.DELETE(":id", l.deleteLeaveForm)
}

// 獲取請假單列表
//
//	route => GET /api/v1/daily/leave
func (l *LeaveApi) getLeaveFormByEmployee(c *gin.Context) {
	var (
		ctx  = gbhttp.Ctx(c)
		in   model.GetLeaveFormByEmployeeReq
		out  model.GetLeaveFormByEmployeeRes
		page = gbhttp.Get(c, "paginator").Interface().(*paginator.Pagination)
		err  error
	)

	in.EmployeeID = gbhttp.Get(c, "employee.id").Uint()

	if err = gbhttp.ParseQuery(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidQuery, err.Error())
		return
	}

	out, err = service.Leave(ctx, page).GetLeaveFormByEmployee(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 刪除請假單
//
//	route => DELETE /api/v1/daily/leave/:id
func (l *LeaveApi) deleteLeaveForm(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.DeleteLeaveFormReq
		err error
	)

	in.ID = gbconv.Uint(c.Param("id"))

	err = service.Leave(ctx).DeleteLeaveForm(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 建立或更新請假單
//
//	route => POST /api/v1/daily/leave
func (l *LeaveApi) saveLeaveForm(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.SaveLeaveFormReq
		out model.SaveLeaveFormRes
		err error
	)

	in.EmployeeID = gbhttp.Get(c, "employee.id").Uint()

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	out, err = service.Leave(ctx).SaveLeaveForm(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}
