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

type LeaveApi struct{}

func (l *LeaveApi) Init(group *gin.RouterGroup) {
	v1 := group.Group("query/leave").Use(middleware.Auth())
	v1.GET("correct", middleware.Web(), l.getLeaveCorrect)
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
