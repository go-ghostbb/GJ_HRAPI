package api

import (
	gberror "ghostbb.io/gb/errors/gb_error"
	gbhttp "ghostbb.io/gb/net/gb_http"
	"github.com/gin-gonic/gin"
	"hrapi/apps/system/v1/model"
	"hrapi/apps/system/v1/service"
	"hrapi/internal/middleware"
	"hrapi/internal/types/enum"
	. "hrapi/internal/utils/response"
	. "hrapi/internal/utils/response/status"
)

type BaseApi struct{}

func (b *BaseApi) Init(v1 *gin.RouterGroup) {
	v1.GET("migrate", b.migrate)
	v1.POST("login/:type", b.login)
	v1.Use(middleware.Auth()).DELETE("logout/:type", b.logout)

	// me
	meGroup := v1.Group("me", middleware.Auth())
	meGroup.GET("info", b.getEmployeeInfo)
	meGroup.GET("perm", b.getPermission)
	meGroup.GET("menu", b.getMenu)
}

// 遷移資料庫
//
//	route => GET /api/v1/migrate
func (b *BaseApi) migrate(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		err error
	)

	err = service.Base(ctx).Migrate()
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}
	Responder(Mount(c)).Ok()
}

// 登入
//
//	route => POST /api/v1/login/:type
func (b *BaseApi) login(c *gin.Context) {
	var (
		ctx       = gbhttp.Ctx(c)
		loginType enum.MenuShow
		in        model.LoginReq
		out       model.LoginRes
		err       error
	)

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	switch gbhttp.ParseParam(c, "type") {
	case "software":
		loginType = enum.Software
	case "web":
		loginType = enum.Web
	default:
		Responder(Mount(c)).Fail(CodeRequestInvalidParam)
		return
	}

	out, err = service.Base(ctx).Login(loginType, in)
	if err != nil {
		if gberror.Is(err, service.ErrUserFrozen) || gberror.Is(err, service.ErrPasswordIncorrect) || gberror.Is(err, service.ErrUserNotFound) {
			Responder(Mount(c)).FailWithMsg(CodeLoginFailed, err.Error())
			return
		}
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 登出
//
//	route => DELETE /api/v1/logout/:type
func (b *BaseApi) logout(c *gin.Context) {
	var (
		ctx        = gbhttp.Ctx(c)
		employeeID = gbhttp.Get(c, "employee.id").Uint()
		logoutType enum.MenuShow
		token      = gbhttp.GetBearerToken(c)
		err        error
	)

	switch gbhttp.ParseParam(c, "type") {
	case "software":
		logoutType = enum.Software
	case "web":
		logoutType = enum.Web
	default:
		Responder(Mount(c)).Fail(CodeRequestInvalidParam)
		return
	}

	err = service.Base(ctx).Logout(logoutType, employeeID, token)
	if err != nil {
		Responder(Mount(c)).Fail(CodeInternalError)
		return
	}

	Responder(Mount(c)).Ok()
}

// 根據token解析出來的employee id獲取employee info
//
//	route => GET /api/v1/me/info
func (b *BaseApi) getEmployeeInfo(c *gin.Context) {
	var (
		ctx        = gbhttp.Ctx(c)
		employeeID = gbhttp.Get(c, "employee.id").Uint()
		out        model.GetEmployeeInfoRes
		err        error
	)

	out, err = service.Base(ctx).GetEmployeeInfo(employeeID)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 根據token解析出來的roles獲取permission
//
//	route => GET /api/v1/me/perm
func (b *BaseApi) getPermission(c *gin.Context) {
	var (
		ctx       = gbhttp.Ctx(c)
		roleCodes = gbhttp.Get(c, "employee.roles").Strings()
		out       []string
		err       error
	)

	out, err = service.Base(ctx).GetPermission(roleCodes)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 根據token解析出來的roles獲取menu
//
//	route => GET /api/v1/me/menu
func (b *BaseApi) getMenu(c *gin.Context) {
	var (
		ctx       = gbhttp.Ctx(c)
		roleCodes = gbhttp.Get(c, "employee.roles").Strings()
		show      = gbhttp.Get(c, "auth.type").Interface().(enum.MenuShow)
		out       []*model.GetMenuRes
		err       error
	)

	out, err = service.Base(ctx).GetMenu(roleCodes, show)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}
