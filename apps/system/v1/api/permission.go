package api

import (
	gbhttp "ghostbb.io/gb/net/gb_http"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"github.com/gin-gonic/gin"
	"hrapi/apps/system/v1/model"
	"hrapi/apps/system/v1/service"
	"hrapi/internal/middleware"
	"hrapi/internal/utils/paginator"

	. "hrapi/internal/utils/response"
	. "hrapi/internal/utils/response/status"
)

type PermissionApi struct{}

func (p *PermissionApi) Init(group *gin.RouterGroup) {
	// 需要有後台權限 (middleware.Software())
	v1 := group.Group("permission").Use(middleware.Auth(), middleware.Software())
	v1.GET("", middleware.Paginator(), p.getByKeyword)
	v1.GET(":id", p.getByID)
	v1.POST("", p.insert)
	v1.PUT(":id", p.update)
	v1.DELETE(":id", p.delete)
	v1.PATCH(":id/status", p.setStatus)
}

// 根據keyword, status獲取權限標示
//
//	route => GET /api/v1/permission
func (p *PermissionApi) getByKeyword(c *gin.Context) {
	var (
		ctx  = gbhttp.Ctx(c)
		in   model.GetByKeywordPermReq
		out  model.GetByKeywordPermRes
		page = gbhttp.Get(c, "paginator").Interface().(*paginator.Pagination)
		err  error
	)

	if err = gbhttp.ParseQuery(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidQuery, err.Error())
		return
	}

	out, err = service.Permission(ctx, page).GetByKeyword(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 根據ID獲取權限標示
//
//	route => GET /api/v1/permission/:id
func (p *PermissionApi) getByID(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.GetByIDPermReq
		out model.GetByIDPermRes
		err error
	)

	in.ID = gbconv.Uint(c.Param("id"))

	out, err = service.Permission(ctx).GetByID(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 新增權限標示
//
//	route => POST /api/v1/permission
func (p *PermissionApi) insert(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.PostPermReq
		err error
	)

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.Permission(ctx).Insert(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 更新權限標示
//
//	route => PUT /api/v1/permission/:id
func (p *PermissionApi) update(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.PutPermReq
		err error
	)

	in.ID = gbconv.Uint(c.Param("id"))

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.Permission(ctx).Update(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 刪除權限標示
//
//	route => DELETE /api/v1/permission/:id
func (p *PermissionApi) delete(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.DeletePermReq
		err error
	)

	in.ID = gbconv.Uint(c.Param("id"))

	err = service.Permission(ctx).Delete(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 更新狀態
//
//	route => PATCH /api/v1/permission/:id/status
func (p *PermissionApi) setStatus(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.SetStatusPermReq
		err error
	)

	in.ID = gbconv.Uint(c.Param("id"))

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.Permission(ctx).SetStatus(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}
