package api

import (
	gbhttp "ghostbb.io/gb/net/gb_http"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"github.com/gin-gonic/gin"
	"hrapi/apps/hr/setting/v1/model"
	"hrapi/apps/hr/setting/v1/service"
	"hrapi/internal/middleware"

	. "hrapi/internal/utils/response"
	. "hrapi/internal/utils/response/status"
)

type SignOffSettingApi struct{}

func (s *SignOffSettingApi) Init(group *gin.RouterGroup) {
	// 需要有後台權限 (middleware.Software())
	v1 := group.Group("signOff").Use(middleware.Auth(), middleware.Software())

	// 請假
	v1.GET("leave", s.getLeaveSignOffSetting)
	v1.PUT("leave/:departmentID/:leaveID/batch", s.updateLeaveBatch)

	// 加班
	v1.GET("overtime", s.getOvertimeSignOffSetting)
	v1.PUT("overtime/:departmentID/:vacationID/batch", s.updateOvertimeBatch)
}

// 查詢請假謙核流程
//
//	route => GET /api/v1/signOff/leave
func (s *SignOffSettingApi) getLeaveSignOffSetting(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.GetLeaveSignOffSettingReq
		out []*model.GetLeaveSignOffSettingRes
		err error
	)

	if err = gbhttp.ParseQuery(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidQuery, err.Error())
		return
	}

	out, err = service.SignOffSetting(ctx).GetLeaveSignOffSetting(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 批量更新leave sign off setting
//
//	route => PUT /api/v1/signOff/leave/:departmentID/:leaveID/batch
func (s *SignOffSettingApi) updateLeaveBatch(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.PutBatchLeaveSignOffSettingReq
		err error
	)

	in.LeaveID = gbconv.Uint(c.Param("leaveID"))
	in.DepartmentID = gbconv.Uint(c.Param("departmentID"))
	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.SignOffSetting(ctx).UpdateLeaveSignOffSettingBatch(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 查詢請假謙核流程
//
//	route => GET /api/v1/signOff/overtime
func (s *SignOffSettingApi) getOvertimeSignOffSetting(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.GetOvertimeSignOffSettingReq
		out []*model.GetOvertimeSignOffSettingRes
		err error
	)

	if err = gbhttp.ParseQuery(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidQuery, err.Error())
		return
	}

	out, err = service.SignOffSetting(ctx).GetOvertimeSignOffSetting(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 批量更新leave sign off setting
//
//	route => PUT /api/v1/signOff/overtime/:departmentID/:vacationID/batch
func (s *SignOffSettingApi) updateOvertimeBatch(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.PutBatchOvertimeSignOffSettingReq
		err error
	)

	in.VacationID = gbconv.Uint(c.Param("vacationID"))
	in.DepartmentID = gbconv.Uint(c.Param("departmentID"))
	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.SignOffSetting(ctx).UpdateOvertimeSignOffSettingBatch(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}
