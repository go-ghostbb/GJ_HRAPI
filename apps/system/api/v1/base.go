package v1

import (
	gberror "ghostbb.io/gb/errors/gb_error"
	gbhttp "ghostbb.io/gb/net/gb_http"
	"github.com/gin-gonic/gin"
	"hrapi/apps/system/model"
	"hrapi/apps/system/service"
	"hrapi/internal/middleware"
	. "hrapi/internal/utils/response"
	. "hrapi/internal/utils/response/status"
)

type BaseApi struct{}

func (b *BaseApi) Init(group *gin.RouterGroup) {
	v1 := group.Group("api/v1")
	v1.GET("migrate", b.migrate)
	v1.POST("login", b.login)
	v1.Use(middleware.Auth()).DELETE("logout", b.logout)
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
		Responder(Mount(c)).FailWithDetail(CodeOperationFailed, err.Error())
		return
	}
	Responder(Mount(c)).Ok()
}

// 登入
//
//	route => POST /api/v1/login
func (b *BaseApi) login(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.LoginReq
		out model.LoginRes
		err error
	)

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
	}

	out, err = service.Base(ctx).Login(in)
	if err != nil {
		if gberror.Is(err, service.ErrUserFrozen) || gberror.Is(err, service.ErrPasswordIncorrect) || gberror.Is(err, service.ErrUserNotFound) {
			Responder(Mount(c)).FailWithMsg(CodeLoginFailed, err.Error())
			return
		}
		Responder(Mount(c)).FailWithDetail(CodeOperationFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 登出
//
//	route => DELETE /api/v1/logout
func (b *BaseApi) logout(c *gin.Context) {
	var (
		ctx        = gbhttp.Ctx(c)
		employeeID = gbhttp.Get(c, "employee.id").Uint()
		token      = gbhttp.GetBearerToken(c)
		err        error
	)

	err = service.Base(ctx).Logout(employeeID, token)
	if err != nil {
		Responder(Mount(c)).Fail(CodeInternalError)
		return
	}

	Responder(Mount(c)).Ok()
}
