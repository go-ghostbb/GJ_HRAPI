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

type RoleApi struct{}

func (r *RoleApi) Init(group *gin.RouterGroup) {
	// 需要有後台權限 (middleware.Software())
	v1 := group.Group("role").Use(middleware.Auth(), middleware.Software())
	v1.GET("", middleware.Paginator(), r.getByKeyword)
	v1.GET(":id", r.getByID)
	v1.POST("", r.insert)
	v1.PUT(":id", r.update)
	v1.DELETE(":id", r.delete)
	v1.PATCH(":id/status", r.setStatus)
	v1.PATCH(":id/menu", r.setMenu)
	v1.PATCH(":id/permission", r.setPerm)
}

// 根據keyword, status獲取角色
//
//	route => GET /api/v1/role
func (r *RoleApi) getByKeyword(c *gin.Context) {
	var (
		ctx  = gbhttp.Ctx(c)
		in   model.GetByKeywordRoleReq
		out  model.GetByKeywordRoleRes
		page = gbhttp.Get(c, "paginator").Interface().(*paginator.Pagination)
		err  error
	)

	if err = gbhttp.ParseQuery(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidQuery, err.Error())
		return
	}

	out, err = service.Role(ctx, page).GetByKeyword(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 根據ID獲取角色
//
//	route => GET /api/v1/role/:id
func (r *RoleApi) getByID(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.GetByIDRoleReq
		out model.GetByIDRoleRes
		err error
	)

	in.ID = gbconv.Uint(c.Param("id"))

	out, err = service.Role(ctx).GetByID(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 新增角色
//
//	route => POST /api/v1/role
func (r *RoleApi) insert(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.PostRoleReq
		err error
	)

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.Role(ctx).Insert(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 更新角色
//
//	route => PUT /api/v1/role/:id
func (r *RoleApi) update(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.PutRoleReq
		err error
	)

	in.ID = gbconv.Uint(c.Param("id"))

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.Role(ctx).Update(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 刪除角色
//
//	route => DELETE /api/v1/role/:id
func (r *RoleApi) delete(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.DeleteRoleReq
		err error
	)

	in.ID = gbconv.Uint(c.Param("id"))

	err = service.Role(ctx).Delete(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 更新狀態
//
//	route => PATCH /api/v1/role/:id/status
func (r *RoleApi) setStatus(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.SetStatusRoleReq
		err error
	)

	in.ID = gbconv.Uint(c.Param("id"))

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.Role(ctx).SetStatus(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 設定可用menu
//
//	route => PATCH /api/v1/role/:id/menu
func (r *RoleApi) setMenu(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.SetMenuRoleReq
		err error
	)

	in.ID = gbconv.Uint(c.Param("id"))

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.Role(ctx).SetMenu(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 設定可用perm
//
//	route => PATCH /api/v1/role/:id/permission
func (r *RoleApi) setPerm(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.SetPermRoleReq
		err error
	)

	in.ID = gbconv.Uint(c.Param("id"))

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.Role(ctx).SetPerm(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}
