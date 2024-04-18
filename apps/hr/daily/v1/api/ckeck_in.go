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

type CheckInApi struct{}

func (ch *CheckInApi) Init(group *gin.RouterGroup) {
	// 需要有網頁端權限 (middleware.Web())
	v1 := group.Group("daily/checkIn").Use(middleware.Auth(), middleware.Web())
	v1.GET("", middleware.Paginator(), ch.getCheckInFormByEmployee)
	v1.POST("", ch.saveCheckInForm)
	v1.DELETE(":id", ch.deleteCheckInForm)
}

// 獲取請假單列表
//
//	route => GET /api/v1/daily/checkIn
func (ch *CheckInApi) getCheckInFormByEmployee(c *gin.Context) {
	var (
		ctx  = gbhttp.Ctx(c)
		in   model.GetCheckInFormByEmployeeReq
		out  model.GetCheckInFormByEmployeeRes
		page = gbhttp.Get(c, "paginator").Interface().(*paginator.Pagination)
		err  error
	)

	in.EmployeeID = gbhttp.Get(c, "employee.id").Uint()

	if err = gbhttp.ParseQuery(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidQuery, err.Error())
		return
	}

	out, err = service.CheckIn(ctx, page).GetCheckInFormByEmployee(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 建立或更新補打卡單
//
//	route => POST /api/v1/daily/checkIn
func (ch *CheckInApi) saveCheckInForm(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.SaveCheckInFormReq
		out model.SaveCheckInFormRes
		err error
	)

	in.EmployeeID = gbhttp.Get(c, "employee.id").Uint()

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	out, err = service.CheckIn(ctx).SaveCheckInForm(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 刪除補打卡單
//
//	route => DELETE /api/v1/daily/checkIn/:id
func (ch *CheckInApi) deleteCheckInForm(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.DeleteCheckInFormReq
		err error
	)

	in.ID = gbconv.Uint(c.Param("id"))

	err = service.CheckIn(ctx).DeleteCheckInForm(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}
