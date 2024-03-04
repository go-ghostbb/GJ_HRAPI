package api

import (
	gbhttp "ghostbb.io/gb/net/gb_http"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"github.com/gin-gonic/gin"
	"hrapi/apps/system/v1/model"
	"hrapi/apps/system/v1/service"
	"hrapi/internal/middleware"

	. "hrapi/internal/utils/response"
	. "hrapi/internal/utils/response/status"
)

type DepartmentApi struct{}

func (d *DepartmentApi) Init(group *gin.RouterGroup) {
	// 需要有後台權限 (middleware.Software())
	v1 := group.Group("department").Use(middleware.Auth(), middleware.Software())
	v1.GET("", d.getByKeyword)
	v1.GET(":id", d.getByID)
	v1.POST("", d.insert)
	v1.PUT(":id", d.update)
	v1.DELETE(":id", d.delete)
	v1.PATCH(":id/status", d.setStatus)
}

// 根據keyword, status獲取部門
//
//	route => GET /api/v1/department
func (d *DepartmentApi) getByKeyword(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.GetByKeywordDepartmentReq
		out []*model.GetByKeywordDepartmentRes
		err error
	)

	if err = gbhttp.ParseQuery(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidQuery, err.Error())
		return
	}

	out, err = service.Department(ctx).GetByKeyword(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 根據ID獲取部門
//
//	route => GET /api/v1/department/:id
func (d *DepartmentApi) getByID(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.GetByIDDepartmentReq
		out model.GetByIDDepartmentRes
		err error
	)

	in.ID = gbconv.Uint(c.Param("id"))

	out, err = service.Department(ctx).GetByID(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 新增部門
// route => POST /api/v1/department
func (d *DepartmentApi) insert(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.PostDepartmentReq
		err error
	)

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.Department(ctx).Insert(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 更新部門
//
//	route => PUT /api/v1/department/:id
func (d *DepartmentApi) update(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.PutDepartmentReq
		err error
	)

	in.ID = gbconv.Uint(c.Param("id"))

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.Department(ctx).Update(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 刪除部門
//
//	route => DELETE /api/v1/department/:id
func (d *DepartmentApi) delete(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.DeleteDepartmentReq
		err error
	)

	in.ID = gbconv.Uint(c.Param("id"))

	err = service.Department(ctx).Delete(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 設置狀態
//
//	route => PATCH /api/v1/department/:id/status
func (d *DepartmentApi) setStatus(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.SetStatusDepartmentReq
		err error
	)

	in.ID = gbconv.Uint(c.Param("id"))

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.Department(ctx).SetStatus(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}
