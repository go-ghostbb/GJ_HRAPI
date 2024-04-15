package api

import (
	gbhttp "ghostbb.io/gb/net/gb_http"
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
	v1.GET("", middleware.Paginator(), ch.getLeaveFormByEmployee)
	v1.POST("", ch.saveLeaveForm)
}

// 獲取請假單列表
//
//	route => GET /api/v1/daily/checkIn
func (ch *CheckInApi) getLeaveFormByEmployee(c *gin.Context) {
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
func (ch *CheckInApi) saveLeaveForm(c *gin.Context) {
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
