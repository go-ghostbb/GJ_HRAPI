package api

import (
	gbhttp "ghostbb.io/gb/net/gb_http"
	"github.com/gin-gonic/gin"
	"hrapi/apps/hr/query/v1/model"
	"hrapi/apps/hr/query/v1/service"
	"hrapi/internal/middleware"

	. "hrapi/internal/utils/response"
	. "hrapi/internal/utils/response/status"
)

type CheckInApi struct{}

func (ch *CheckInApi) Init(group *gin.RouterGroup) {
	// 需要有網頁端權限 (middleware.Web())
	v1 := group.Group("query/checkIn").Use(middleware.Auth(), middleware.Web())
	v1.GET("status", ch.queryStatus)
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
