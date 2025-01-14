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

type LeaveApi struct{}

func (l *LeaveApi) Init(group *gin.RouterGroup) {
	// 需要有後台權限 (middleware.Software())
	v1 := group.Group("leave").Use(middleware.Auth())
	v1.GET("", middleware.Paginator(), l.getByKeyword)
	v1.GET(":id", middleware.Software(), l.getByID)
	v1.POST("", middleware.Software(), l.insert)
	v1.PUT(":id", middleware.Software(), l.update)
	v1.DELETE(":id", middleware.Software(), l.delete)
	v1.PATCH(":id/status", middleware.Software(), l.setStatus)
	v1.POST("reset/available", middleware.Software(), l.resetEmployeeAvailable)
	// 群組操作
	v1.GET("group/:leaveID", middleware.Software(), l.getLeaveGroup)
	v1.POST("group", middleware.Software(), l.insertGroup)
	v1.DELETE("group/:id", middleware.Software(), l.deleteGroup)
	v1.PATCH("group/:id/name", middleware.Software(), l.setLeaveGroupName)
	v1.PATCH("group/:id/employee", middleware.Software(), l.setLeaveGroupEmployee)
	v1.PATCH("group/:id/condition", middleware.Software(), l.setLeaveGroupCond)
}

// 根據keyword, status獲取對應假別
//
//	route => /api/v1/leave
func (l *LeaveApi) getByKeyword(c *gin.Context) {
	var (
		ctx  = gbhttp.Ctx(c)
		in   model.GetByKeywordLeaveReq
		out  model.GetByKeywordLeaveRes
		page = gbhttp.Get(c, "paginator").Interface().(*paginator.Pagination)
		err  error
	)

	if err = gbhttp.ParseQuery(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidQuery, err.Error())
		return
	}

	out, err = service.Leave(ctx, page).GetByKeyword(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 根據ID獲取對應假別
//
//	route => /api/v1/leave/:id
func (l *LeaveApi) getByID(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.GetByIDLeaveReq
		out model.GetByIDLeaveRes
		err error
	)
	in.ID = gbconv.Uint(c.Param("id"))

	out, err = service.Leave(ctx).GetByID(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 新增假別
//
//	route => /api/v1/leave
func (l *LeaveApi) insert(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.PostLeaveReq
		err error
	)
	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.Leave(ctx).Insert(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 更新假別
//
//	route => /api/v1/leave/:id
func (l *LeaveApi) update(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.PutLeaveReq
		err error
	)
	in.ID = gbconv.Uint(c.Param("id"))

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.Leave(ctx).Update(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 刪除假別
//
//	route => /api/v1/leave/:id
func (l *LeaveApi) delete(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.DeleteLeaveReq
		err error
	)
	in.ID = gbconv.Uint(c.Param("id"))

	err = service.Leave(ctx).Delete(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 設置狀態
//
//	route => PATCH /api/v1/leave/:id
func (l *LeaveApi) setStatus(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.SetStatusLeaveReq
		err error
	)
	in.ID = gbconv.Uint(c.Param("id"))
	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.Leave(ctx).SetStatus(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 獲取假別的群組
//
//	route => GET /api/v1/leave/group/:leaveID
func (l *LeaveApi) getLeaveGroup(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.GetLeaveGroupReq
		out []*model.GetLeaveGroupRes
		err error
	)
	in.LeaveID = gbconv.Uint(c.Param("leaveID"))

	out, err = service.Leave(ctx).GetLeaveGroup(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 新增群組
//
//	route => POST /api/v1/leave/group
func (l *LeaveApi) insertGroup(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.PostLeaveGroupReq
		err error
	)

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.Leave(ctx).InsertGroup(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 刪除群組
//
//	route => DELETE /api/v1/leave/group/:id
func (l *LeaveApi) deleteGroup(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.DeleteLeaveGroupReq
		err error
	)
	in.ID = gbconv.Uint(c.Param("id"))

	err = service.Leave(ctx).DeleteGroup(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 設置群組名稱
//
//	route => PATCH /api/v1/leave/group/:id/name
func (l *LeaveApi) setLeaveGroupName(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.SetLeaveGroupNameReq
		err error
	)
	in.ID = gbconv.Uint(c.Param("id"))
	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.Leave(ctx).SetLeaveGroupName(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 設置群組的員工
//
//	route => PATCH /api/v1/leave/group/:id/employee
func (l *LeaveApi) setLeaveGroupEmployee(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.SetLeaveGroupEmployeeReq
		err error
	)
	in.ID = gbconv.Uint(c.Param("id"))

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.Leave(ctx).SetLeaveGroupEmployee(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 設置群組的條件
//
//	route => PATCH /api/v1/leave/group/:id/condition
func (l *LeaveApi) setLeaveGroupCond(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.SetLeaveGroupCondReq
		err error
	)
	in.ID = gbconv.Uint(c.Param("id"))

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.Leave(ctx).SetLeaveGroupCond(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 重新計算請假可用天數
//
//	route => POST /api/v1/leave/reset/available
func (l *LeaveApi) resetEmployeeAvailable(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.ResetEmployeeAvailableReq
		err error
	)

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.Leave(ctx).ResetEmployeeAvailable(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}
