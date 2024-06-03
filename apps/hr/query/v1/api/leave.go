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

type LeaveApi struct{}

func (l *LeaveApi) Init(group *gin.RouterGroup) {
	v1 := group.Group("query/leave").Use(middleware.Auth())
	v1.GET("correct", middleware.Web(), l.getLeaveCorrect)
	v1.GET("correct/keyword", middleware.Software(), middleware.Paginator(), l.getLeaveCorrectKeyword)

	v1S := group.Group("leaveCorrect").Use(middleware.Auth())
	v1S.PUT(":id", middleware.Software(), l.updateLeaveCorrect)
}

// 查詢可用假別
//
//	route => GET /api/v1/query/leave/correct
func (l *LeaveApi) getLeaveCorrect(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.GetLeaveCorrectReq
		out []*model.GetLeaveCorrectRes
		err error
	)

	in.EmployeeID = gbhttp.Get(c, "employee.id").Uint()

	if err = gbhttp.ParseQuery(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidQuery, err.Error())
		return
	}

	out, err = service.Leave(ctx).GetLeaveCorrect(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 根據關鍵字查詢可用假別
//
//	route => /api/v1/query/leave/correct/keyword
func (l *LeaveApi) getLeaveCorrectKeyword(c *gin.Context) {
	var (
		ctx  = gbhttp.Ctx(c)
		page = gbhttp.Get(c, "paginator").Interface().(*paginator.Pagination)
		in   model.GetLeaveCorrectByKeywordReq
		out  model.GetLeaveCorrectByKeywordRes
		err  error
	)

	if err = gbhttp.ParseQuery(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidQuery, err.Error())
		return
	}

	out, err = service.Leave(ctx, page).GetLeaveCorrectByKeyword(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 更新leave correct
//
//	route => PUT /api/v1/leaveCorrect/:id
func (l *LeaveApi) updateLeaveCorrect(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.PutLeaveCorrectReq
		err error
	)

	in.ID = gbconv.Uint(c.Param("id"))

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.Leave(ctx).UpdateLeaveCorrect(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}
