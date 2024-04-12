package api

import (
	gbhttp "ghostbb.io/gb/net/gb_http"
	"github.com/gin-gonic/gin"
	"hrapi/apps/hr/sign-off/v1/model"
	"hrapi/apps/hr/sign-off/v1/service"
	"hrapi/internal/types"

	. "hrapi/internal/utils/response"
	. "hrapi/internal/utils/response/status"
)

type LeaveApi struct{}

func (l *LeaveApi) Init(group *gin.RouterGroup) {
	v1 := group.Group("sign-off/leave")
	v1.GET("", l.getByUUID)
	v1.PUT("approve/:uuid", l.approve)
	v1.PUT("reject/:uuid", l.reject)
}

// 根據uuid獲取請假單資訊
//
//	route => GET /api/v1/sign-off/leave
func (l *LeaveApi) getByUUID(c *gin.Context) {
	var (
		ctx  = gbhttp.Ctx(c)
		uuid string
		out  *types.LeaveRequestForm
		err  error
	)
	uuid = c.Query("uuid")

	if uuid == "" {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidQuery, "uuid is required")
		return
	}

	out, err = service.Leave(ctx).GetByUUID(uuid)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 核准
//
//	route => PUT /api/v1/sign-off/leave/approve
func (l *LeaveApi) approve(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.LeaveApproveReq
		err error
	)

	in.UUID = c.Param("uuid")
	if in.UUID == "" {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidQuery, "uuid is required")
		return
	}

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.Leave(ctx).Approve(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 駁回
//
//	route => PUT /api/v1/sign-off/leave/reject
func (l *LeaveApi) reject(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.LeaveRejectReq
		err error
	)

	in.UUID = c.Param("uuid")
	if in.UUID == "" {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidQuery, "uuid is required")
		return
	}

	in.UUID = c.Param("uuid")
	if in.UUID == "" {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidQuery, "uuid is required")
		return
	}

	err = service.Leave(ctx).Reject(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}
