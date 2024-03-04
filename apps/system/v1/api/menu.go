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

	. "hrapi/internal/utils/response"
	. "hrapi/internal/utils/response/status"
)

type MenuApi struct{}

func (m *MenuApi) Init(group *gin.RouterGroup) {
	// 需要有後台權限 (middleware.Software())
	v1 := group.Group("menu").Use(middleware.Auth(), middleware.Software())
	v1.GET("", m.getByKeyword)
	v1.GET(":id", m.getByID)
	v1.POST("", m.insert)
	v1.PUT(":id", m.update)
	v1.DELETE(":id", m.delete)
	v1.PATCH(":id/status", m.setStatus)
}

// 根據keyword, status, show獲取對應menu
//
//	route => GET /api/v1/menu
func (m *MenuApi) getByKeyword(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.GetByKeywordMenuReq
		out []*model.GetByKeywordMenuRes
		err error
	)
	if err = gbhttp.ParseQuery(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidQuery, err.Error())
		return
	}

	out, err = service.Menu(ctx).GetByKeyword(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 根據ID獲取menu
//
//	route => GET /api/v1/menu/:id
func (m *MenuApi) getByID(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.GetByIDMenuReq
		out model.GetByIDMenuRes
		err error
	)

	in.ID = gbconv.Uint(c.Param("id"))

	out, err = service.Menu(ctx).GetByID(in)
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

// 新增menu
//
//	route => POST /api/v1/menu
func (m *MenuApi) insert(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.PostMenuReq
		err error
	)

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.Menu(ctx).Insert(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 更新menu
//
//	route => PUT /api/v1/menu/:id
func (m *MenuApi) update(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.PutMenuReq
		err error
	)
	in.ID = gbconv.Uint(c.Param("id"))

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.Menu(ctx).Update(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 刪除menu
//
//	route => DELETE /api/v1/menu/:id
func (m *MenuApi) delete(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.DeleteMenuReq
		err error
	)
	in.ID = gbconv.Uint(c.Param("id"))

	err = service.Menu(ctx).Delete(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 設置menu狀態
//
//	route => PATCH /api/v1/menu/:id/status
func (m *MenuApi) setStatus(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.SetStatusMenuReq
		err error
	)
	in.ID = gbconv.Uint(c.Param("id"))

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.Menu(ctx).SetStatus(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}
