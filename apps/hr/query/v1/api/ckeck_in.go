package api

import (
	gbhttp "ghostbb.io/gb/net/gb_http"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"github.com/gin-gonic/gin"
	"hrapi/apps/hr/query/v1/model"
	"hrapi/apps/hr/query/v1/service"
	"hrapi/internal/middleware"
	"hrapi/internal/utils/paginator"

	. "hrapi/internal/utils/response"
	. "hrapi/internal/utils/response/status"
)

type CheckInApi struct{}

func (ch *CheckInApi) Init(group *gin.RouterGroup) {
	// 需要有網頁端權限 (middleware.Web())
	v1 := group.Group("query/checkIn").Use(middleware.Auth())
	v1.GET("status", middleware.Web(), ch.queryStatus)
	v1.GET("status/keyword", middleware.Software(), middleware.Paginator(), ch.queryStatusByKeyword)

	v1S := group.Group("checkIn").Use(middleware.Auth())
	v1S.PUT(":id", middleware.Software(), ch.updateCheckInStatus)
}

// 查詢出勤狀態
//
//	route => GET /api/v1/query/checkIn/status
func (ch *CheckInApi) queryStatus(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.QueryStatusReq
		out []*model.QueryStatusRes
		err error
	)

	in.EmployeeID = gbhttp.Get(c, "employee.id").Uint()

	if err = gbhttp.ParseQuery(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidQuery, err.Error())
		return
	}

	out, err = service.CheckIn(ctx).QueryStatus(in)
	if err != nil {
		Responder(Mount(c)).FailWithDetail(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 查詢出勤狀態(關鍵字)
//
//	route => GET /api/v1/query/checkIn/status/keyword
func (ch *CheckInApi) queryStatusByKeyword(c *gin.Context) {
	var (
		ctx  = gbhttp.Ctx(c)
		page = gbhttp.Get(c, "paginator").Interface().(*paginator.Pagination)
		in   model.QueryStatusByKeywordReq
		out  model.QueryStatusByKeywordRes
		err  error
	)

	if err = gbhttp.ParseQuery(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidQuery, err.Error())
		return
	}

	out, err = service.CheckIn(ctx, page).QueryStatusByKeyword(in)
	if err != nil {
		Responder(Mount(c)).FailWithDetail(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 更新狀態表
//
//	route => PUT /api/v1/checkIn/:id
func (ch *CheckInApi) updateCheckInStatus(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.PutCheckInStatusReq
		err error
	)

	in.ID = gbconv.Uint(c.Param("id"))

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.CheckIn(ctx).UpdateCheckInStatus(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}
