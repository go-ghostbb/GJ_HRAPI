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

type VacationApi struct{}

func (v *VacationApi) Init(group *gin.RouterGroup) {
	// 需要有後台權限 (middleware.Software())
	v1 := group.Group("vacation").Use(middleware.Auth(), middleware.Software())
	v1.GET("", middleware.Paginator(), v.getByKeyword)
	v1.GET(":id", v.getByID)
	v1.POST("", v.insert)
	v1.PUT(":id", v.update)
	v1.DELETE(":id", v.delete)
	v1.PATCH(":id/status", v.setStatus)
	// 群組操作
	v1.GET("group/:vacationID", v.getVacationGroup)
	v1.POST("group", v.insertGroup)
	v1.DELETE("group/:id", v.deleteGroup)
	v1.PATCH("group/:id/employee", v.setVacationGroupEmployee)
	v1.PATCH("group/:id/condition", v.setVacationGroupCond)
}

// 根據keyword, status獲取對應休假日
//
//	route => /api/v1/vacation
func (v *VacationApi) getByKeyword(c *gin.Context) {
	var (
		ctx  = gbhttp.Ctx(c)
		in   model.GetByKeywordVacationReq
		out  model.GetByKeywordVacationRes
		page = gbhttp.Get(c, "paginator").Interface().(*paginator.Pagination)
		err  error
	)

	if err = gbhttp.ParseQuery(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidQuery, err.Error())
		return
	}

	out, err = service.Vacation(ctx, page).GetByKeyword(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 根據ID獲取對應休假日
//
//	route => /api/v1/vacation/:id
func (v *VacationApi) getByID(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.GetByIDVacationReq
		out model.GetByIDVacationRes
		err error
	)
	in.ID = gbconv.Uint(c.Param("id"))

	out, err = service.Vacation(ctx).GetByID(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 新增休假日
//
//	route => /api/v1/vacation
func (v *VacationApi) insert(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.PostVacationReq
		err error
	)
	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.Vacation(ctx).Insert(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 更新休假日
//
//	route => /api/v1/vacation/:id
func (v *VacationApi) update(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.PutVacationReq
		err error
	)
	in.ID = gbconv.Uint(c.Param("id"))

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.Vacation(ctx).Update(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 刪除休假日
//
//	route => /api/v1/vacation/:id
func (v *VacationApi) delete(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.DeleteVacationReq
		err error
	)
	in.ID = gbconv.Uint(c.Param("id"))

	err = service.Vacation(ctx).Delete(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 設置狀態
//
//	route => PATCH /api/v1/vacation/:id
func (v *VacationApi) setStatus(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.SetStatusVacationReq
		err error
	)
	in.ID = gbconv.Uint(c.Param("id"))
	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.Vacation(ctx).SetStatus(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 獲取休假日的群組
//
//	route => GET /api/v1/leave/group/:vacationID
func (v *VacationApi) getVacationGroup(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.GetVacationGroupReq
		out []*model.GetVacationGroupRes
		err error
	)
	in.VacationID = gbconv.Uint(c.Param("vacationID"))

	out, err = service.Vacation(ctx).GetVacationGroup(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 新增群組
//
//	route => POST /api/v1/vacation/group
func (v *VacationApi) insertGroup(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.PostVacationGroupReq
		err error
	)

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.Vacation(ctx).InsertGroup(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 刪除群組
//
//	route => DELETE /api/v1/leave/group/:id
func (v *VacationApi) deleteGroup(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.DeleteVacationGroupReq
		err error
	)
	in.ID = gbconv.Uint(c.Param("id"))

	err = service.Vacation(ctx).DeleteGroup(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 設置群組的員工
//
//	route => PATCH /api/v1/leave/group/:id/employee
func (v *VacationApi) setVacationGroupEmployee(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.SetVacationGroupEmployeeReq
		err error
	)
	in.ID = gbconv.Uint(c.Param("id"))

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.Vacation(ctx).SetVacationGroupEmployee(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 設置群組的條件
//
//	route => PATCH /api/v1/leave/group/:id/overtimeRate
func (v *VacationApi) setVacationGroupCond(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.SetVacationGroupOvertimeRateReq
		err error
	)
	in.ID = gbconv.Uint(c.Param("id"))

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.Vacation(ctx).SetVacationGroupOvertimeRate(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}
