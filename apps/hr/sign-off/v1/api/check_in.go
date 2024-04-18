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

type CheckInApi struct{}

func (ch *CheckInApi) Init(group *gin.RouterGroup) {
	v1 := group.Group("sign-off/checkIn")
	v1.GET("", ch.getByUUID)
	v1.PUT("approve/:uuid", ch.approve)
	v1.PUT("reject/:uuid", ch.reject)
}

// 根據uuid獲取補打卡單資訊
//
//	route => GET /api/v1/sign-off/checkIn
func (ch *CheckInApi) getByUUID(c *gin.Context) {
	var (
		ctx  = gbhttp.Ctx(c)
		uuid string
		out  *types.CheckInRequestForm
		err  error
	)
	uuid = c.Query("uuid")

	if uuid == "" {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidQuery, "uuid is required")
		return
	}

	out, err = service.CheckIn(ctx).GetByUUID(uuid)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 核准
//
//	route => PUT /api/v1/sign-off/checkIn/approve
func (ch *CheckInApi) approve(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.CheckInApproveReq
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

	err = service.CheckIn(ctx).Approve(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 駁回
//
//	route => PUT /api/v1/sign-off/checkIn/reject
func (ch *CheckInApi) reject(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.CheckInRejectReq
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

	err = service.CheckIn(ctx).Reject(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}
