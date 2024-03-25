package api

import (
	gbhttp "ghostbb.io/gb/net/gb_http"
	"github.com/gin-gonic/gin"
	"hrapi/apps/hr/filing/v1/model"
	"hrapi/apps/hr/filing/v1/service"
	"hrapi/internal/middleware"
	. "hrapi/internal/utils/response"
	. "hrapi/internal/utils/response/status"
)

type CheckInStatusApi struct{}

func (s *CheckInStatusApi) Init(group *gin.RouterGroup) {
	// 需要有後台權限 (middleware.Software())
	v1 := group.Group("checkInStatus").Use(middleware.Auth(), middleware.Software())
	v1.POST("filing", s.filing)
	v1.POST("upload", s.uploadData)
}

// 輸入的日期區間, 將班表寫進打卡狀態裡
//
//	route => POST /api/v1/checkInStatus/filing
func (s *CheckInStatusApi) filing(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.FilingCheckInStatusReq
		err error
	)

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.CheckInStatus(ctx).Filing(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 上傳資料
//
//	route => POST /api/v1/checkInStatus/upload
func (s *CheckInStatusApi) uploadData(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  []*model.UploadDataReq
		err error
	)

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.CheckInStatus(ctx).UploadData(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}
