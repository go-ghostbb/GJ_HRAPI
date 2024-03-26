package api

import (
	gberror "ghostbb.io/gb/errors/gb_error"
	gbhttp "ghostbb.io/gb/net/gb_http"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"hrapi/apps/system/v1/model"
	"hrapi/apps/system/v1/service"
	"hrapi/internal/middleware"
	"hrapi/internal/utils/paginator"
	. "hrapi/internal/utils/response"
	. "hrapi/internal/utils/response/status"
)

type EmployeeApi struct{}

func (e *EmployeeApi) Init(group *gin.RouterGroup) {
	// 需要有後台權限 (middleware.Software())
	v1 := group.Group("employee").Use(middleware.Auth(), middleware.Software())
	v1.GET("", middleware.Paginator(), e.getByKeyword)
	v1.GET(":id", e.getByID)
	v1.POST("", e.insert)
	v1.PUT(":id", e.update)
	v1.DELETE(":id", e.delete)
	v1.PATCH(":id/password", e.resetPassword)
	v1.PATCH(":id/employmentStatus", e.setEmploymentStatus)
	v1.PATCH(":id/login/status", e.setLoginStatus)
	v1.GET(":id/checkInStatus", e.getByDateRangeCheckInStatus)
}

// 根據keyword獲取employee資料
//
//	route => GET /api/v1/employee?keyword=test
func (e *EmployeeApi) getByKeyword(c *gin.Context) {
	var (
		ctx  = gbhttp.Ctx(c)
		in   model.GetByKeywordEmployeeReq
		out  model.GetByKeywordEmployeeRes
		page = gbhttp.Get(c, "paginator").Interface().(*paginator.Pagination)
		err  error
	)
	if err = gbhttp.ParseQuery(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidQuery, err.Error())
		return
	}

	out, err = service.Employee(ctx, page).GetByKeyword(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 根據id獲取employee資料
//
//	route => GET /api/v1/employee/:id
func (e *EmployeeApi) getByID(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.GetByIDEmployeeReq
		out model.GetByIDEmployeeRes
		err error
	)
	in.ID = gbconv.Uint(gbhttp.ParseParam(c, "id"))

	out, err = service.Employee(ctx).GetByID(in)
	if err != nil {
		if gberror.Is(err, gorm.ErrRecordNotFound) {
			Responder(Mount(c)).OkWithDetail("record not found")
			return
		}
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 新增一筆employee資料
//
//	route => POST /api/v1/employee
func (e *EmployeeApi) insert(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.PostEmployeeReq
		out model.PostEmployeeRes
		err error
	)
	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	out, err = service.Employee(ctx).Insert(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 更新employee資料
//
//	route => PUT /api/v1/employee/:id
func (e *EmployeeApi) update(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.PutEmployeeReq
		err error
	)
	in.ID = gbconv.Uint(c.Query("id"))
	if in.ID = gbconv.Uint(c.Param("id")); in.ID == 0 {
		Responder(Mount(c)).FailWithMsg(CodeRequestInvalidParam, err.Error())
		return
	}

	if err = gbhttp.ParseJSON(c, &in.Employee); err != nil {
		Responder(Mount(c)).FailWithMsg(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.Employee(ctx).Update(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 根據id刪除employee資料
//
//	route => DELETE /api/v1/employee/:id
func (e *EmployeeApi) delete(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.DeleteEmployeeReq
		err error
	)

	in.ID = gbconv.Uint(gbhttp.ParseParam(c, "id"))

	err = service.Employee(ctx).Delete(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 重置密碼
//
//	route => PATCH /api/v1/employee/:id/password
func (e *EmployeeApi) resetPassword(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.ResetPasswordReq
		err error
	)

	in.ID = gbconv.Uint(gbhttp.ParseParam(c, "id"))

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.Employee(ctx).ResetPassword(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 設置就職狀態
//
//	route => PATCH /api/v1/employee/:id/employmentStatus
func (e *EmployeeApi) setEmploymentStatus(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.SetEmploymentStatusReq
		err error
	)
	in.ID = gbconv.Uint(c.Param("id"))

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidQuery, err.Error())
		return
	}

	err = service.Employee(ctx).SetEmploymentStatus(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 設置登入狀態, 啟用, 禁用
//
//	route => PATCH /api/v1/employee/:id/login/status
func (e *EmployeeApi) setLoginStatus(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.SetLoginStatusReq
		err error
	)
	in.ID = gbconv.Uint(c.Param("id"))

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidQuery, err.Error())
		return
	}

	err = service.Employee(ctx).SetLoginStatus(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 獲取員工打卡狀態
//
//	route => GET /api/v1/employee/:id/checkInStatus
func (e *EmployeeApi) getByDateRangeCheckInStatus(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.GetByDateRangeCheckInStatusReq
		out []model.GetByDateRangeCheckInStatusRes
		err error
	)

	in.EmployeeID = gbconv.Uint(c.Param("id"))

	if err = gbhttp.ParseQuery(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidQuery, err.Error())
		return
	}

	out, err = service.Employee(ctx).GetByDateRangeCheckInStatus(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}
